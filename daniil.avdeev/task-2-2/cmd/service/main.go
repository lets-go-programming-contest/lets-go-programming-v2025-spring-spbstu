package main

import (
	"container/heap"
	"fmt"

	"github.com/realFrogboy/task-2-2/internal/parser"
)

func main() {
	n_positions, error := parser.GetInt()
	if error != nil {
		fmt.Println(error)
		return
	}
	if n_positions < 0 {
		fmt.Printf("invalid number of menu positions: %d\n", n_positions)
		return
	}

	pq, error := parser.GetMenuPositions(n_positions)
	if error != nil {
		fmt.Println(error)
		return
	}

	kth_max, error := parser.GetInt()
	if error != nil {
		fmt.Println(error)
		return
	}
	if kth_max < 0 || kth_max > pq.Len() {
		fmt.Printf("invalid kth maximum position: %d\n", kth_max)
		return
	}

	var res int
	for cnt := 0; cnt < kth_max; cnt++ {
		res = (*pq)[0]
		heap.Pop(pq)
	}

	fmt.Println(res)
}
