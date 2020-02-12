package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/akamensky/argparse"
	"github.com/santosh/gophercises/link"
)

var siteFlag *string
var depthFlag *int

var sitemapTmpl = `
<?xml version="1.0" encoding="UTF-8"?>

<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
{{ range . }}
   <url>
      <loc>{{ . }}</loc>
   </url>
{{ end }}
</urlset>
`

func init() {
	parser := argparse.NewParser("sitemap", "Generates sitemaps of given website.")

	siteFlag = parser.String("s", "site", &argparse.Options{Help: "Site to be crawled", Default: "http://127.0.0.1:6060/"})
	depthFlag = parser.Int("d", "depth", &argparse.Options{Help: "Max depth to go starting from --site", Default: 2})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}
}

// No other protocols than https.
func main() {
	link := bfs(*siteFlag, *depthFlag)

	t := template.Must(template.New("").Parse(sitemapTmpl))
	err := t.Execute(os.Stdout, link)
	if err != nil {
		panic(err)
	}
}

func bfs(urlStr string, maxDepth int) []string {
	// Refrain going into visited links to prevent cyclic condition.
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}

			seen[url] = struct{}{}
			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nq[link] = struct{}{}
				}
			}
		}
	}
	var ret []string
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}
	resp, err := client.Get(urlStr) // GET urlStr
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	// Use URL from the response, eliminates HTTP redirect cases
	reqURL := resp.Request.URL

	// eliminates the case of picking up entire path
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(base, hrefs(resp.Body, base))
}

func filter(base string, links []string) []string {
	var ret []string
	for _, link := range links {
		// Bye bye external links
		if !strings.HasPrefix(link, base) {
			continue
		}

		// No fragments (e.g.`#links`); only the page they are on
		if strings.Contains(link, "#") {
			canonical := strings.Split(link, "#")[0]
			ret = append(ret, canonical)
			continue
		}

		if strings.HasPrefix(link, base) {
			ret = append(ret, link)
		}
	}

	return ret
}

// TODO: Parse concurrently for efficiency.
func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r) // search for links

	var ret []string
	for _, l := range links {
		if l.Href == "" {
			// we don't link blank hrefs
			continue
		}

		switch {
		// Relative paths will be converted to absolute URL
		case strings.HasPrefix(l.Href, "/") && !strings.HasPrefix(l.Href, "//"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		default:
		}
	}

	return ret
}
