package main

import "fmt"

// Variables to override via build flags
var (
	Version    = "unset"
	BuildTime  = "unset"
	CommitHash = "unset"
)

func main() {
	printBuildInfo()
	fmt.Println("Application logic...")
}

func printBuildInfo() {
	fmt.Println("=== Build Information ===")
	fmt.Printf("Version:    %s\n", Version)
	fmt.Printf("Build Time: %s\n", BuildTime)
	fmt.Printf("Git Commit: %s\n", CommitHash)
	fmt.Println("=========================")
}
