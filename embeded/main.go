package main

import _ "embed"

//go:embed hello.txt
var s string

func main() {
	print(s)
}
