package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// interface for mocking
type FileSystem interface {
	getMarkdownFileList() []string
	renameFile(from, to string)
	readFile(name string) []byte
}

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
		err := os.Rename(from, to)
		checkFatal(err)
	}
}

// return file contents as array of bytes
func (*realFileSystem) readFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	checkFatal(err)
	return data
}

// get list of files in directory
func getFileListInDir(dirName string) []string {
	dir, err := os.Open(dirName)
	checkFatal(err)
	fileNames, err := dir.Readdirnames(-1)
	checkFatal(err)
	return fileNames
}

// read metadata from pandoc yaml header file
func readFileMetadata(fileName string) *metadata {
	meta := &metadata{}
	yaml.Unmarshal(fs.readFile(fileName), meta)
	return meta
}

// return file contents as string, discard commented lines
func readOptionsFile(name string) string {
	file, err := os.Open(name)
	checkFatal(err)
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cl := scanner.Text()
		if !strings.HasPrefix(cl, "#") {
			lines = append(lines, cl)
		}
	}
	checkFatal(scanner.Err())
	return strings.Join(lines, "\n")
}

// read all files into a single string
func mergeFilesToBuffer() string {
	inputFiles := fs.getMarkdownFileList()
	var buffer []byte
	for _, path := range inputFiles {
		data := fs.readFile(path)
		buffer = append(buffer, data...)
	}
	return string(buffer)
}
