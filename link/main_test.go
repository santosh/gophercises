package link

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

var htmlFile io.Reader

func TestParse(t *testing.T) {
	htmlFile = strings.NewReader(`<html>
	<body>
	  <h1>Hello!</h1>
	  <a href="/other-page">A link to another page</a>
	  <a href="/second-page">A link to second page</a>
	</body>
	
	</html>`)

	expected := []Link{
		{Href: "/other-page", Text: "A link to another page"},
		{Href: "/second-page", Text: "A link to second page"},
	}

	got, _ := Parse(htmlFile)

	if !reflect.DeepEqual(got, expected) {
		t.Error("Parse not working correctly.")
	}

}
