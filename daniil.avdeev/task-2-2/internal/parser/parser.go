package parser

import (
	"container/heap"
	"fmt"

	"github.com/realFrogboy/task-2-2/internal/priority_queue"
)

func GetInt() (int, error) {
	var val int

	_, error := fmt.Scanf("%d", &val)
	if error != nil {
		return 0, error
	}

	return val, nil
}

func GetMenuPositions(n int) (*priority_queue.IntHeap, error) {
	priorityQueue := &priority_queue.IntHeap{}
	heap.Init(priorityQueue)

	for cnt := 0; cnt < n; cnt++ {
		rating, error := GetInt()
		if error != nil {
			return priorityQueue, error
		}

		heap.Push(priorityQueue, rating)
	}

	return priorityQueue, nil
}
