package main

import "testing"

func TestImageCheck(t *testing.T) {
	imgCheckDir = "test/img"
	_, ok := allImagesUsed()
	if ok {
		t.Error("Not all images are used: should have complained")
	}
}
