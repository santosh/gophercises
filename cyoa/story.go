package cyoa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>Choose Your Own Adventure</title>
</head>

<body>
  <section  class="page">
    <h1>{{ .Title }}</h1>
    {{range .Story}}
    <p>{{ . }}</p>
    {{end}}

    <ul>
        {{range .Options}}
        	<li><a href="/{{ .Chapter }}">{{.Text}}</a></li>
        {{end}}
	</ul>
  </section>
</body>

</html>`

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s Stories, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s      Stories
	t      *template.Template
	pathFn func(r *http.Request) string
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:] // by omiting the '/' we get the arc
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)

		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong..", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

// JSONStory takes filename and returns Stories
func JSONStory(fileFlag string) (Stories, error) {
	var stories Stories
	jsonContent, err := ioutil.ReadFile(fileFlag)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonContent, &stories)
	return stories, err
}

// Stories is a map of Chapters
type Stories map[string]Chapter

type options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// Chapter is used to Unmarshal provided json
type Chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []options `json:"options"`
}
