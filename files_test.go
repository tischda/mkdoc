package main

import (
	"reflect"
	"testing"
)

func TestReadOptionsFile(t *testing.T) {
	expected := `
--from=markdown+yaml_metadata_block

--highlight-style=tango
--number-sections
--top-level-division=chapter
--listings
--filter pandoc-citeproc
--standalone
--pdf-engine-opt=--miktex-admin
--variable=papersize:a4paper
--variable=geometry:margin=1in
--variable=date={{.Tag}}~gen.~{{.Date}}~-~{{.Time}}

--variable=day={{.Day}}
--variable=month={{.Month}}
--variable=year={{.Year}}

--variable=hour={{.Hour}}
--variable=minute={{.Minute}}
--variable=second={{.Second}}

--to=latex+table_captions

-o {{.Target}}`

	actual := readOptionsFile("test/pandoc.options")
	checkEquals(t, expected, actual)
}

func TestGetImageFileList(t *testing.T) {
	actual := getFileListInDir("test/img")
	expected := []string{
		"test-img-1.png",
		"test-img-2.png",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %q, but was: %q", expected, actual)
	}
}
