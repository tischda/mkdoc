package main

import (
	"reflect"
	"testing"
)

func TestCompileNewNames(t *testing.T) {

	// standard case
	testCase(t, map[string]string{
		"04-aa": "01-aa",
		"07-bb": "02-bb",
		"08-cc": "03-cc"})

	// upward remaning 01-bb would overwrite 02-bb
	testCase(t, map[string]string{
		"01-aa": "01-aa",
		"01-bb": "02-bb",
		"02-bb": "03-bb"})

	// downward renaming 03-aa would overwrite 02-aa
	testCase(t, map[string]string{
		"02-aa": "01-aa",
		"03-aa": "02-aa",
		"03-bb": "03-bb"})

	// both up-downward renaming
	testCase(t, map[string]string{
		"00-aa": "01-aa",
		"01-aa": "02-aa",
		"02-aa": "03-aa",
		"05-bb": "04-bb",
		"06-bb": "05-bb"})
}

func testCase(t *testing.T, expected map[string]string) {
	mfs := &mockFileSystem{getSortedKeys(expected)}
	fs = mfs
	renumberFiles()

	if !reflect.DeepEqual(mfs.fileList, getSortedValues(expected)) {
		t.Errorf("Expected: %q, was: %q", expected, mfs.fileList)
	}
}

