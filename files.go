package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// get list of files in directory
func getFileListInDir(dirName string) []string {
	dir, err := os.Open(dirName)
	if err != nil {
		log.Fatalln(err)
	}
	defer dir.Close()
	fileNames, err := dir.Readdirnames(-1)
	if err != nil {
		log.Fatalln(err)
	}
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
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cl := scanner.Text()
		if !strings.HasPrefix(cl, "#") {
			lines = append(lines, cl)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatalln(err)
	}
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
