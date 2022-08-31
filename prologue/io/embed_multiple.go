package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed sample.json version.txt

var static embed.FS

func main() {
	b, err := static.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(b))

	b2, err := fs.ReadFile(static, "version.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("version: %s\n", string(b2))
}