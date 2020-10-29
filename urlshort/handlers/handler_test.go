package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/santosh/gophercises/urlshort/handlers"
)

var expectedPairs = []handlers.Pair{
	handlers.Pair{
		Path: "/yaml",
		URL:  "https://en.wikipedia.org/wiki/YAML",
	},
	handlers.Pair{
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

	parsedYAML, err := handlers.ParseYAML(yamlContent)

	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(parsedYAML, expectedPairs) {
		t.Errorf("Wrong YAML Parsing.")
	}

}

func TestMapHandler(t *testing.T) {
	pathsToURLs := map[string]string{
		"/gh": "https://github.com/santosh",
	}

	t.Run("test with path available", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/gh", nil)
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(handlers.MapHandler(pathsToURLs, nil))

		handler.ServeHTTP(rr, req)

		// Check if status code represent a redirection
		if status := rr.Code; status != http.StatusFound {
			t.Errorf("handler showig wrong status code: get %v want %v", status, http.StatusFound)
		}
	})
	t.Run("test with path unavailable", func(t *testing.T) {

	})
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

	parsedJSON, err := handlers.ParseJSON(jsonContent)
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(parsedJSON, expectedPairs) {
		t.Errorf("Wrong JSON Parsing. Got \n%s\n want \n%s\n", parsedJSON, expectedPairs)
	}
}

// Test for buildMap - should return path map (which MapHandler accepts) from yaml data
func TestBuildMap(t *testing.T) {
	pathsToUrls := handlers.BuildMap(expectedPairs)

	if !reflect.DeepEqual(pathsToUrls, expectedMap) {
		t.Errorf("Problem during map generation.")
	}
}
