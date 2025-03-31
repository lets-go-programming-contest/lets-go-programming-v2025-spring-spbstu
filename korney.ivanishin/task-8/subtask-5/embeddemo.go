package main

import "embed"

//go:embed data/single_file.txt
var fileString string

//go:embed data/single_file.txt
var fileByte []byte

//go:embed data/*.hash
var folder embed.FS

func main() {
        print(fileString)
        print(string(fileByte))
        content1, _ := folder.ReadFile("data/file1.hash")
        print(string(content1))
        content2, _ := folder.ReadFile("data/file2.hash")
        print(string(content2))
}
