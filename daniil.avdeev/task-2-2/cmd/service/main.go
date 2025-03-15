package main

import (
	"container/heap"
	"fmt"

	"github.com/realFrogboy/task-2-2/internal/parser"
)

func main() {
	nPositions, error := parser.GetInt()
	if error != nil {
		fmt.Println(error)
		return
	}
	if nPositions < 0 {
		fmt.Printf("invalid number of menu positions: %d\n", nPositions)
		return
	}

	priorityQueue, error := parser.GetMenuPositions(nPositions)
	if error != nil {
		fmt.Println(error)
		return
	}

	kthMax, error := parser.GetInt()
	if error != nil {
		fmt.Println(error)
		return
	}
	if kthMax < 0 || kthMax > priorityQueue.Len() {
		fmt.Printf("invalid kth maximum position: %d\n", kthMax)
		return
	}

	var res int
	for cnt := 0; cnt < kthMax; cnt++ {
		res = (*priorityQueue)[0]
		heap.Pop(priorityQueue)
	}

	fmt.Println(res)
}
