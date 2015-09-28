package main

import (
	"os"
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
	actual := getMarkdownInputFiles()
	checkDeepEquals(t, expected, actual)
}
