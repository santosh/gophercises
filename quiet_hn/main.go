package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/santosh/gophercises/quiet_hn/hn"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := getTopStories(numStories)
		if err != nil {
			// it is essential to keep in mind what kind of information are
			// being split outside our application. You probably don't want
			// to give out information like what kind of database you are using
			// to a potential attacker.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func getTopStories(numStories int) ([]item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("Failed to load top stories")
	}
	// we need something to come out of goroutine,
	// it will be this struct
	type result struct {
		idx  int
		item item
		err  error
	}
	resultCh := make(chan result)

	for i := 0; i < numStories; i++ {
		// pass into to the goroutine error or the item
		go func(idx, id int) {
			hnItem, err := client.GetItem(id)
			if err != nil {
				resultCh <- result{idx: idx, err: err}
			}
			resultCh <- result{idx: idx, item: parseHNItem(hnItem)}
		}(i, ids[i])
	}

	var results []result

	for i := 0; i < numStories; i++ {
		results = append(results, <-resultCh)
	}

	// now sort the slice as there is no gurantee all goroutines
	// would have returned in same order they were executed
	sort.Slice(results, func(i, j int) bool {
		return results[i].idx < results[j].idx
	})

	// separate faulty items and pass healthy one to the caller
	var stories []item
	for _, res := range results {
		if res.err != nil {
			continue
		}
		// only the story type is considered
		if isStoryLink(res.item) {
			stories = append(stories, res.item)
		}
	}

	return stories, nil
}

// isStoryLink is used to filter out posts other than 'story' type
func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

// parseHNItem parses a "Item" resource from hackernews and
// returns a local "Item" with "www." trimmed.
func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

// templateData data structure is sent to template to populate the page
type templateData struct {
	Stories []item
	Time    time.Duration
}
