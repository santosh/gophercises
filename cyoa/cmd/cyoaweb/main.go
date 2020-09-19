package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/akamensky/argparse"
	"github.com/santosh/gophercises/cyoa"
)

var (
	fileFlag *string
	portFlag *int
)

func init() {
	parser := argparse.NewParser("cyoaweb", "Web version of The Little Blue Gopher. A choose your own adventure type game.")
	fileFlag = parser.String("f", "file", &argparse.Options{Required: true, Help: "File to use for story."})
	portFlag = parser.Int("p", "port", &argparse.Options{Required: false, Help: "Port to start the CYOA webapp on."})

	err := parser.Parse(os.Args)

	if err != nil && *fileFlag == "" {
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	// if no port is pasesd, default to 3030
	if *portFlag == 0 {
		*portFlag = 3030
	}
}

func main() {
	stories, err := cyoa.JSONStory(*fileFlag)
	if err != nil {
		exit("JSON parsing failed.")
	}

	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := cyoa.NewHandler(stories,
		cyoa.WithPathFunc(pathFn),
		cyoa.WithTemplate(tpl))

	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting server on http://127.0.0.1:%d\n", *portFlag)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), h))

}

func exit(msg string) {
	log.Println(msg)
	os.Exit(1)
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):] // by omiting the '/' we get the arc
}

var storyTmpl = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>Choose Your Own Adventure</title>
	<style>
	  body {font-family: helvetica, arial; }
	  h1 {text-align: center; position: relative; }
	  .page {
		  width: 80%;
		  max-width: 500px;
		  margin: auto;
		  margin-top: 40px;
		  margin-bottom: 40px;
		  padding: 80px;
		  background: #fffcf6;
		  border: 1px solid #eee;
		  box-shadow: 0 10px 6px -6px #777;
	  }
	  ul {
		  border-top: 1px dotted #ccc;
		  padding: 10px 0 0 0;
		  -webkit-padding-start: 0;
	  }
	  li {
		  padding-top: 10px;
	  }
	  a,
	  a:visited {
		  text-decoration: none;
		  color: #6295b5;
	  }
	  a:active,
	  a:hover {
		  color: #7792a2;
	  }
	  p {
		  text-indent: 1em;
	  }
	</style>
</head>

<body>
	  <section  class="page">
    <h1>{{ .Title }}</h1>
    {{range .Story}}
    <p>{{ . }}</p>
    {{end}}

    <ul>
        {{range .Options}}
        	<li><a href="/story/{{ .Chapter }}">{{.Text}}</a></li>
        {{end}}
	</ul>
	</section>
</body>

</html>`
