package main

import (
	"embed"
	"os"
	"text/template"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
var f embed.FS

//go:embed template/*.tmpl
var templates embed.FS

func main() {
	print(s)
	print(string(b))

	data, _ := f.ReadFile("hello.txt")
	print(string(data))

	t, err := template.ParseFS(templates, "template/*")
	if err != nil {
		panic(err)
	}

	if err := t.ExecuteTemplate(os.Stdout, "example.tmpl", "Gopher"); err != nil {
		panic(err)
	}
}
