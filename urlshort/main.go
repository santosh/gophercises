package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/akamensky/argparse"
)

func getContent(fp string) (yamlContent []byte) {
	if fp == "" {
		fp = "urls.yaml"
	}

	yamlContent, err := ioutil.ReadFile(fp)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the YAML file: %s\n", fp))
	}
	return
}

var yamlFlag *string

func init() {
	parser := argparse.NewParser("urlshort", "URL Shortener with multiple data backstores.")
	yamlFlag = parser.String("y", "yaml", &argparse.Options{Required: false, Help: "YAML file to be used"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/ghme":           "https://github.com/santosh",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	yamlContent := getContent(*yamlFlag)

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
