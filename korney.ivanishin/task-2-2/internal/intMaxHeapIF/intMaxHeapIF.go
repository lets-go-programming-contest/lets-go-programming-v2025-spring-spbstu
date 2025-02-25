package intMaxHeapIF

type IntMaxHeap []int32

func (h IntMaxHeap) Len() int {
	return len(h)
}

func (h IntMaxHeap) Less(idx1 int, idx2 int) bool {
	return h[idx1] > h[idx2]
}

func (h IntMaxHeap) Swap(idx1 int, idx2 int) {
	h[idx1], h[idx2] = h[idx2], h[idx1]
}

func (h *IntMaxHeap) Push(val any) {
	*h = append(*h, val.(int32))
}

func (h *IntMaxHeap) Pop() any {
	oldHeap := *h
	oldHeapLen := len(oldHeap)
	val := oldHeap[oldHeapLen - 1]
	*h = oldHeap[0 : oldHeapLen - 1]
	return val
}
