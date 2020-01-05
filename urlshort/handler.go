package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler returns an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathsToUrls[r.URL.Path]; ok {
			log.Println("Matched:", path)
			http.Redirect(w, r, path, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

// Pair maps pair of Path and URL to be Unmarshalled
type Pair struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

// YAMLHandler parses the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pairs []Pair
	err := yaml.Unmarshal(yml, &pairs)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string, len(pairs))

	for _, entry := range pairs {
		pathsToUrls[entry.Path] = entry.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}

// JSONHandler parses the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
// 	{
// 		"path": "/example",
// 		"url": "https://www.example.com"
// 	}
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsonContent []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pairs []Pair
	err := json.Unmarshal(jsonContent, &pairs)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string, len(pairs))

	for _, entry := range pairs {
		pathsToUrls[entry.Path] = entry.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}
