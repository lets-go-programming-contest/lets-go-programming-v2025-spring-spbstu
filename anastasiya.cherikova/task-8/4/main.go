package main

//go:generate sh -c "echo 'package main\n\nconst Version = \"1.2.3\"' > version.go"

import "fmt"

func main() {
	fmt.Println("Current Version:", Version)
	fmt.Println("Build Time:", BuildTime)
}
