package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// interface for mocking
type FileSystem interface {
	getMarkdownFileList() []string
	renameFile(from, to string)
	readFile(name string) []byte
}

// marker for real file system
type realFileSystem struct{}

var fs FileSystem = &realFileSystem{}

// get sorted list of Markdown input files from current directory
func (*realFileSystem) getMarkdownFileList() []string {
	fileList, _ := filepath.Glob("[0-9][0-9]*.md")
	sort.Strings(fileList)
	return fileList
}

// rename file
func (*realFileSystem) renameFile(from, to string) {
	if from != to {
		fmt.Printf("  %s --> %s\n", from, to)
		if err := os.Rename(from, to); err != nil {
			log.Fatalln(err)
		}
	}
}

// return file contents as array of bytes
func (*realFileSystem) readFile(name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
