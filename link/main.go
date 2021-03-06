package link

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Link stores "href" attribute as well as text within the <a> tag
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return
// a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

// buildLink takes a node and returns a Link struct
func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

// text pulls text from given node
func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Main is disabled here, see next lesson in the queue.
// func main() {
// 	data, err := os.Open(os.Args[1])
// 	if err != nil {
// 		exit("Error opening the html file.")
// 	}

// 	doc, err := Parse(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(doc)
// }
