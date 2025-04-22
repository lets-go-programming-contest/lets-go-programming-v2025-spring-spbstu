package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

)

//go:embed data/string.txt
var fileString string

//go:embed data/number.txt
var fileNumber string  

func main() {
	numberStr := strings.TrimSpace(fileNumber)

	num, err := strconv.Atoi(numberStr)
	if err != nil {
		panic(fmt.Sprintf("Conversion error: %v\nFile content: %q", err, fileNumber))
	}

	fmt.Println("String:", fileString)
	fmt.Println("Num:", num)
	fmt.Println("NumNum:", num*num)
}

