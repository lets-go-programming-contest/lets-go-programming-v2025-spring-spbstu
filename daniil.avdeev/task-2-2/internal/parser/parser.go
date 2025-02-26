package parser

import (
	"container/heap"
	"fmt"

	"github.com/realFrogboy/task-2-2/internal/priority_queue"
)

func GetInt() (int, error) {
	var Val int

	_, error := fmt.Scanf("%d", &Val)
	if error != nil {
		return 0, error
	}

	return Val, nil
}

func GetMenuPositions(n int) (*priority_queue.IntHeap, error) {
	h := &priority_queue.IntHeap{}
	heap.Init(h)

	for cnt := 0; cnt < n; cnt++ {
		rating, error := GetInt()
		if error != nil {
			return h, error
		}

		heap.Push(h, rating)
	}

	return h, nil
}
