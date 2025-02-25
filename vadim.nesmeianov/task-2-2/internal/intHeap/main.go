package IntHeap

import (
	"container/heap"
	"errors"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // The sorting order changed
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Get k by order in heap
func (h IntHeap) GetKNode(k int) (int, error) {
	if h.Len() < k {
		errStr := fmt.Sprintf("too big k = %d for heap with len = %d", k, h.Len())
		return 0, errors.New(errStr)
	}

	heap.Init(&h)
	var last int
	for i := 0; i < k; i += 1 {
		last = heap.Pop(&h).(int)
	}
	return last, nil
}
