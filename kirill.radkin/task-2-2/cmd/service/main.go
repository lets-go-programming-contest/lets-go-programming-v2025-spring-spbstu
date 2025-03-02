package main

import (
	myheap "github.com/yanelox/task-2-2/internal/IntHeap"

	"container/heap"
	"fmt"
)

func main() {
	h := &myheap.IntHeap{}
	heap.Init(h)

	var numDishes int
	_, err := fmt.Scan(&numDishes)

	if err != nil {
		fmt.Println("Incorrect input")
		return
	}

	for i := 0; i < numDishes; i++ {
		var dishRating int

		_, err := fmt.Scan(&dishRating)

		if err != nil {
			fmt.Println("Incorrect input")
			return
		}

		heap.Push(h, dishRating)
	}

	var choose int
	_, err = fmt.Scan(&choose)

	if err != nil {
		fmt.Println("Incorrect input")
		return
	}

	if choose > numDishes {
		fmt.Printf("Choosed number should be less or equal than number of dishes: %d\n", numDishes)
		return
	}

	for i := 0; i < choose-1; i++ {
		heap.Pop(h)
	}

	fmt.Println(heap.Pop(h))
}
