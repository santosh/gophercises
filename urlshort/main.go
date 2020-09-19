package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/akamensky/argparse"
	"github.com/boltdb/bolt"
	"github.com/santosh/gophercises/urlshort/handlers"
)

func getYAMLContent(fp string) (yamlContent []byte) {
	if fp == "" {
		fp = "urls.yaml"
	}

	yamlContent, err := ioutil.ReadFile(fp)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the YAML file: %s\n", fp))
	}
	return
}

func getJSONContent(fp string) (jsonContent []byte) {
	if fp == "" {
		fp = "urls.json"
	}

	jsonContent, err := ioutil.ReadFile(fp)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the JSON file: %s\n", fp))
	}
	return
}

var yamlFlag, jsonFlag, boltFlag *string

func init() {
	parser := argparse.NewParser("urlshort", "URL Shortener with multiple data backstores.")

	yamlFlag = parser.String("y", "yaml", &argparse.Options{Required: false, Help: "YAML file to be used"})
	jsonFlag = parser.String("j", "json", &argparse.Options{Required: false, Help: "JSON file to be used"})
	boltFlag = parser.String("b", "bolt", &argparse.Options{Required: false, Help: "BoltDB file to be used"})

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
		"/gh":             "https://github.com/santosh",
	}

	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	yamlContent := getYAMLContent(*yamlFlag)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := handlers.YAMLHandler([]byte(yamlContent), mapHandler)
	if err != nil {
		panic(err)
	}

	jsonContent := getJSONContent(*jsonFlag)

	// Build the JSONHandler using the yamlHandler as the fallback
	jsonHandler, err := handlers.JSONHandler([]byte(jsonContent), yamlHandler)
	if err != nil {
		panic(err)
	}

	if *boltFlag == "" {
		*boltFlag = "urls.db"
	}

	db, err := bolt.Open(*boltFlag, 0644, nil)
	if err != nil {
		fmt.Println("Starting the server without BoltDB backend on http://localhost:8080")
		http.ListenAndServe(":8080", jsonHandler)
	}
	defer db.Close()

	boltHandler := handlers.BoltHandler(db, jsonHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server with BoltDB backend on http://localhost:8080")
	http.ListenAndServe(":8080", boltHandler)
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
