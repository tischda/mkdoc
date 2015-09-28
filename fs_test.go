package main

import (
	"os"
	"testing"
)

// assign mock file system
func setup() {
	fs = &mockFileSystem{[]string{"00-first.md", "01-first.md", "02-second.md"}}
}

// define setup (once for all tests)
func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestGetMarkdownFileList(t *testing.T) {
	expected := []string{
		"00-first.md",
		"01-first.md",
		"02-second.md",
	}
	actual := fs.getMarkdownFileList()
	checkDeepEquals(t, expected, actual)
}

func TestMergeFilesToBuffer(t *testing.T) {
	actual := mergeFilesToBuffer()
	expected := "datadatadata"
	if actual != expected {
		t.Errorf("Expected: %q, but was: %q", expected, actual)
	}
}
