package main

import "testing"

func TestReadOptionsFile(t *testing.T) {
	expected := `
--from=markdown+yaml_metadata_block
--listings
--number-sections
--variable=papersize:a4paper
--variable=geometry:margin=1in

--variable=date={{.Tag}}~gen.~{{.Date}}~-~{{.Time}}

-o out/{{.Target}}`

	actual := readOptionsFile("pandoc.options")
	checkEquals(t, expected, actual)
}

func TestMarkdownInputFiles(t *testing.T) {
	expected := []string{
		"01-first.md",
		"02-second.md",
	}
	actual := getMarkdownInputFiles()
	checkDeepEquals(t, expected, actual)
}
