package main

import (
	"embed"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/*.hash
var folder embed.FS

func main() {
	println(fileString)
	println(string(fileByte))
	content1, _ := folder.ReadFile("folder/file1.hash")
	println(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	println(string(content2))
}
