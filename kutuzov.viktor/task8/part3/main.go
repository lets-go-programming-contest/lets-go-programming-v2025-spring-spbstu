package main

import "fmt"

var Version = "development"

func main() {
	fmt.Println("Version: " + Version)
}

//go build -ldflags="-X 'main.Version=v1.0.0'
