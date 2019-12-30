package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO: Keep provided yaml as fallback.

func getContent(fp string) (yamlContent []byte) {
	if fp != "" {
		yamlContent, err := ioutil.ReadFile(fp)
		if err != nil {
			exit(fmt.Sprintf("Failed to open the YAML file: %s\n", fp))
		}
		return yamlContent
	}
	return
}

func main() {
	yamlFilePtr := flag.String("yaml", "", "yaml file to be used")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/ghme":           "https://github.com/santosh",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	yamlContent := getContent(*yamlFilePtr)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := YAMLHandler([]byte(yamlContent), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on http://localhost:8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
