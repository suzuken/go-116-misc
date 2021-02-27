package main

import (
	"fmt"
	"io/fs"
	"os"
)

type fooFS struct{}

func (fsys *fooFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func NewFS() fs.FS {
	return &fooFS{}
}

func main() {
	f := NewFS()
	file, err := f.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("s = %+v\n", s)
}
