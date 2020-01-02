package main

import (
	"reflect"
	"testing"
)

var expectedPairs = []Pair{
	Pair{
		Path: "/yaml",
		URL:  "https://en.wikipedia.org/wiki/YAML",
	},
	Pair{
		Path: "/json",
		URL:  "https://en.wikipedia.org/wiki/JSON",
	},
}

var expectedMap = map[string]string{
	"/yaml": "https://en.wikipedia.org/wiki/YAML",
	"/json": "https://en.wikipedia.org/wiki/JSON",
}

// Test for parseYAML - Should load YAML data
func TestParseYAML(t *testing.T) {
	yamlContent := []byte(`
  - path: /yaml
    url: https://en.wikipedia.org/wiki/YAML
  - path: /json
    url: https://en.wikipedia.org/wiki/JSON`)

	parsedYAML, err := parseYAML(yamlContent)

	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(parsedYAML, expectedPairs) {
		t.Errorf("Wrong YAML Parsing.")
	}

}

// Test for parseJSON - Should load JSON data
func TestParseJSON(t *testing.T) {
	jsonContent := []byte(`[
    {
        "path": "/yaml",
        "url": "https://en.wikipedia.org/wiki/YAML"
    },
    {
        "path": "/json",
        "url": "https://en.wikipedia.org/wiki/JSON"
    }
]`)

	parsedJSON, err := parseJSON(jsonContent)
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(parsedJSON, expectedPairs) {
		t.Errorf("Wrong JSON Parsing. Got \n%s\n want \n%s\n", parsedJSON, expectedPairs)
	}
}

// Test for buildMap - should return path map (which MapHandler accepts) from yaml data
func TestBuildMap(t *testing.T) {
	pathsToUrls := buidMap(expectedPairs)

	if !reflect.DeepEqual(pathsToUrls, expectedMap) {
		t.Errorf("Problem during map generation.")
	}
}
