//go:build !(lin || win)

package main

import "fmt"

func GetOsInfo() {
	fmt.Println("Unknown OS")
}
