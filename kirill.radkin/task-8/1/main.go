package main

var authors = []string {
	"Pushkin",
}

var MyFavouritePockemon = "Bulbasaur"

func main() {
	println("Hello, world!")

	for _, a := range authors {
		println(a)
	}

	println("My favourite pockemon is", MyFavouritePockemon)
}