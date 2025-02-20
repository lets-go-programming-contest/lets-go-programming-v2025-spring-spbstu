package calculate

import (
	"container/heap"

	"github.com/vktr-ktzv/task2-2/internal/intHeap"
)

func FindKthLargest(nums []int, k int) int {
	h := &intHeap.IntHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
