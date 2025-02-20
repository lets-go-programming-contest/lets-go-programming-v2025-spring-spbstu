package main

import (
	"fmt"
	"log"

	"github.com/vktr-ktzv/task2-2/internal/calculate"
	"github.com/vktr-ktzv/task2-2/internal/intReader"
	"github.com/vktr-ktzv/task2-2/internal/sliceReader"
)

func main() {
	var N, k int
	fmt.Println("Ведите количество блюд:")
	N, err := intReader.Read()
	if err != nil {
		log.Fatal(err)
	} else if N <= 0 {
		log.Fatal("error: N <= 0")
	}

	dishes, err := sliceReader.Read(uint(N))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Введите какое блюдо выбираем:")
	k, err = intReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	result := calculate.FindKthLargest(dishes, k)

	fmt.Println(result)
}
