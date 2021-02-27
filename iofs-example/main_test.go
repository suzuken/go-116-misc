package main

import (
	"testing"
	"testing/fstest"
)

func TestFunction(t *testing.T) {
	fs := make(fstest.MapFS)
	fs["file/exists"] = &fstest.MapFile{
		Data: []byte("test"),
	}
	if err := fstest.TestFS(fs, "file/exists"); err != nil {
		t.Fatal(err)
	}
}
