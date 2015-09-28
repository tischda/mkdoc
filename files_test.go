package main

import (
	"os"
	"reflect"
	"testing"
)

func TestReadOptionsFile(t *testing.T) {
	expected := `
--from=markdown+yaml_metadata_block
--listings
--number-sections
--variable=papersize:a4paper
--variable=geometry:margin=1in

--variable=date={{.Tag}}~gen.~{{.Date}}~-~{{.Time}}

-o {{.Target}}`

	actual := readOptionsFile("test/pandoc.options")
	checkEquals(t, expected, actual)
}

func TestMarkdownInputFiles(t *testing.T) {
	expected := []string{
		"00-first.md",
		"01-first.md",
		"02-second.md",
	}
	os.Chdir("test")
	actual := getMarkdownFileList()
	os.Chdir("..")
	checkDeepEquals(t, expected, actual)
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

func TestFillBuffer(t *testing.T) {
	os.Chdir("test")
	actual := mergeFilesToBuffer()
	os.Chdir("..")
	expected := "![no label](img/test-img-2.png)\ndon't overwrite me!\njust text\n"

	if actual != expected {
		t.Errorf("Expected: %q, but was: %q", expected, actual)
	}
}
