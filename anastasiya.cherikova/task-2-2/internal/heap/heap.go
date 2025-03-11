// heap/heap.go
package heap

import "container/heap"

// IntHeap - минимальная куча для хранения целых чисел
type IntHeap []int

// Реализация интерфейса heap.Interface
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Инициализация кучи
func InitHeap(h *IntHeap) {
	heap.Init(h)
}

// Операции с кучей
func PushHeap(h *IntHeap, x int) {
	heap.Push(h, x)
}

func PopHeap(h *IntHeap) int {
	return heap.Pop(h).(int)
}
