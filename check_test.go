package main

import "testing"

func TestCheckImageDirectory(t *testing.T) {
	_, ok := allImagesUsed("test/img")
	if ok {
		t.Error("Not all images are used: should have complained")
	}
}
