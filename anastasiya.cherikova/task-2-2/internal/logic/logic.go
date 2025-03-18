// logic/logic.go
package logic

import (
	"task-2-2/internal/heap"
)

// FindKthLargest находит k-й наибольший элемент
func FindKthLargest(nums []int, k int) int {
	h := &heap.IntHeap{}
	heap.InitHeap(h)

	for _, num := range nums {
		if h.Len() < k {
			heap.PushHeap(h, num)
		} else if num > (*h)[0] {
			heap.PopHeap(h)
			heap.PushHeap(h, num)
		}
	}
	return (*h)[0]
}
