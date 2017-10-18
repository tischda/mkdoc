package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

const dateFormat = "%d.%02d.%02d~-~%02d:%02d:%02d"
const metadataFileName = `metadata.yaml`
const optionsFileName = `pandoc.options`

type metadata struct {
	Target string `yaml:"target"`
	Tag    string
	Date   string
	Day    string
	Month  string
	Year   string
	Time   string
	Hour   string
	Minute string
	Second string
}

// read options and run pandoc on 99-filename.md files
func runPandoc() {
	pandocOptions := readOptionsFile(optionsFileName)
	options := strings.Fields(fillMeta(pandocOptions))

	inputFiles := fs.getMarkdownFileList()
	inputFiles = append(inputFiles, metadataFileName)
	options = append(options, inputFiles...)

	executeProcess("pandoc", options...)
}

// fill pandoc options template with document metadata
// 		Target: obtained from 'metadata.yaml'
// 		Tag: 	'git describe --tags'
// 		Date: 	current date
// 		Time: 	current time
func fillMeta(template string) string {
	meta := readFileMetadata(metadataFileName)

	// make sure target path exists
	path := filepath.Dir(meta.Target)
	if path != "." {
		if _, err := os.Stat(path); os.IsNotExist(err) && !noop {
			fmt.Println("Creating target path:", path)
			checkFatal(os.MkdirAll(path, os.ModePerm))
		}
	}

	meta.Tag = getGitTag()

	t := time.Now()
	meta.Date = formatDate(t)
	meta.Time = formatTime(t)

	meta.Day = fmt.Sprintf("%02d", t.Day())
	meta.Month = fmt.Sprintf("%02d", t.Month())
	meta.Year = fmt.Sprintf("%02d", t.Year())

	meta.Hour = fmt.Sprintf("%02d", t.Hour())
	meta.Minute = fmt.Sprintf("%02d", t.Minute())
	meta.Second = fmt.Sprintf("%02d", t.Second())

	return fillTemplate(template, meta)
}

// return git tag from project directory
func getGitTag() string {
	cmdOut, _ := exec.Command("git", "describe", "--tags").Output()
	return strings.TrimSpace(string(cmdOut))
}

// return formatted date
func formatDate(t time.Time) string {
	return fmt.Sprintf("%02d.%02d.%d", t.Day(), t.Month(), t.Year())
}

// return formatted time
func formatTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

// fill text template with data
func fillTemplate(text string, data interface{}) string {
	template, err := template.New("test").Parse(text)
	checkFatal(err)

	buff := bytes.NewBufferString("")
	err = template.Execute(buff, data)
	checkFatal(err)

	return buff.String()
}
