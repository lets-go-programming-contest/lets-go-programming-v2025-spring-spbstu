package main

import (
	"fmt"

	myheap "github.com/yanelox/task-2-2/pkg/IntHeap"

	"container/heap"
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

	for i := 0; i < choose-1; i++ {
		heap.Pop(h)
	}

	fmt.Println(heap.Pop(h))
}
