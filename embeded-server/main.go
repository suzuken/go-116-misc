package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static
var static embed.FS

func main() {
	fsys, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(fsys)))
	http.ListenAndServe(":8080", nil)
}
