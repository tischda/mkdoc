package main

import (
	"reflect"
	"testing"
)

func checkEquals(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("Expected: %q, but was: %q", expected, actual)
	}
}

func checkDeepEquals(t *testing.T, expected []string, actual []string) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %q, but was: %q", expected, actual)
	}
}
