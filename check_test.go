package main

import "testing"

func TestCheckImageDirectory(t *testing.T) {
	imgageDir = "test/img"
	_, ok := allImagesUsed()
	if ok {
		t.Error("Not all images are used: should have complained")
	}
}
