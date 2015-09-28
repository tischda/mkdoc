package main

import (
	"fmt"
	"os"
	"testing"
)

type mockFileSystem struct {
	fileList []string
}

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

// get sorted list of Markdown input files from current directory
func (fs *mockFileSystem) getMarkdownFileList() []string {
	return fs.fileList
}

// rename file
func (fs *mockFileSystem) renameFile(from, to string) {
	if from != to {
		fmt.Printf("  %s --> %s\n", from, to)
		for i, name := range fs.fileList {
			if name == from {
				fs.fileList[i] = to
			}
		}
	}
}

// return file contents as array of bytes
func (*mockFileSystem) readFile(name string) []byte {
	return []byte("data")
}
