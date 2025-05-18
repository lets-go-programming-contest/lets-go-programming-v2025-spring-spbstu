//go:build !(lin || win || darwin)

package main

import "fmt"

func GetOsInfo() {
	fmt.Println("Unknown OS")
}
