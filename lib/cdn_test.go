package lib

import (
	"bytes"
	testdir "github.com/philpearl/assets/test"
	"html/template"
	"testing"
)

func TestFindVersionedFiles(t *testing.T) {
	cdn := NewCDN("http://hatstand.com/")

	err := cdn.FindVersionedFiles(testdir.FileSystem, testdir.Paths)
	if err != nil {
		t.Fatalf("Error looking for versioned files. %v", err)
	}

	assertName := func(name, exp string) {
		n := cdn.StaticFileName(name)
		if n != exp {
			t.Errorf("For %s - Expected %s, got %s", name, exp, n)
		}
	}
	assertName("static/js/a1.js", "http://hatstand.com/static/js/a1-78b86493ac2e9e54b60471852919ac10.js")
	assertName("static/js/behive.min.js", "http://hatstand.com/static/js/behive.min-8c17f4c608db7de5fef12c54bbdf7783.js")

	funcmap := template.FuncMap{
		"static": cdn.StaticFileName,
	}

	temp, err := template.New("test").Funcs(funcmap).Parse(`<script src="{{static "static/js/a1.js"}}"></script>`)

	var buf bytes.Buffer
	err = temp.Execute(&buf, nil)

	if err != nil {
		t.Fatalf("Failed to run template. %v", err)
	}

	if buf.String() != `<script src="http://hatstand.com/static/js/a1-78b86493ac2e9e54b60471852919ac10.js"></script>` {
		t.Fatalf("template output dodgy. %s", buf.String())
	}
}
