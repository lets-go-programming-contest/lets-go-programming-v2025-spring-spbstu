package main

import "embed"

var authors = []string {
	"Pushkin",
}

var MyFavouritePockemon = "Bulbasaur"

//go:embed test/*.txt
var folder embed.FS

func main() {
	println("Hello, world!")

	for _, a := range authors {
		println(a)
	}

	println("My favourite pockemon is", MyFavouritePockemon)
	println()

	content1, _ := folder.ReadFile("test/tmp1.txt")
	println(string(content1))

	content2, _ := folder.ReadFile("test/tmp2.txt")
	println(string(content2))
}