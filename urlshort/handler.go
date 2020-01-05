package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
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

func buidMap(pairs []Pair) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, entry := range pairs {
		pathsToUrls[entry.Path] = entry.URL
	}
	return pathsToUrls
}

func parseYAML(yamlBytes []byte) ([]Pair, error) {
	var pairs []Pair

	err := yaml.Unmarshal(yamlBytes, &pairs)
	if err != nil {
		return nil, err
	}
	return pairs, nil
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
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yamlBytes)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buidMap(parsedYAML)
	return MapHandler(pathsToUrls, fallback), nil
}

func parseJSON(jsonBytes []byte) ([]Pair, error) {
	var pairs []Pair

	err := json.Unmarshal(jsonBytes, &pairs)
	if err != nil {
		return nil, err
	}
	return pairs, nil
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
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(jsonBytes)
	if err != nil {
		return nil, err
	}
	pathsToUrls := buidMap(parsedJSON)

	return MapHandler(pathsToUrls, fallback), nil
}

var redirects = []byte("redirects")

// BoltHandler pulls out key, vaule from the passed BoltDB file.
func BoltHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var url string
		err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(redirects)
			if bucket == nil {
				return fmt.Errorf("Bucket %q not found", redirects)
			}

			value := bucket.Get([]byte(r.URL.Path))
			if value != nil {
				url = string(value)
			}

			return nil
		})

		if err == nil && url != "" {
			log.Println("Matched:", url)
			http.Redirect(w, r, url, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}
