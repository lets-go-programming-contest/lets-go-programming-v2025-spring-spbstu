package preference

import "container/heap"

// https://pkg.go.dev/container/heap

type PreferenceHeap []int

func (h PreferenceHeap) Len() int {
	return len(h)
}

func (h PreferenceHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h PreferenceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PreferenceHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *PreferenceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *PreferenceHeap) Add(val int) {
	heap.Push(h, val)
}

func GetKthLargest(preferences []int, k int) int {
	h := &PreferenceHeap{}
	heap.Init(h)

	for _, pref := range preferences {
		h.Add(pref)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return heap.Pop(h).(int)
}
