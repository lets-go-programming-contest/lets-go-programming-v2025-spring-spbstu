package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Generate data")
	err := os.WriteFile("../../internal/data/hello.txt", []byte("Hello"), 0644)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("../../internal/data/world.txt", []byte("World!"), 0644)
	if err != nil {
		panic(err)
	}
}
