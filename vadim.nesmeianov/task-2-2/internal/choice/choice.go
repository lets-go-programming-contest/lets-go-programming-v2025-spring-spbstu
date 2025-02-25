package choice

import (
	"container/heap"
	"fmt"
	IntHeap "task-2-2/internal/intHeap"
)

func Run() error {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return err
	}

	h := make(IntHeap.IntHeap, n)
	heap.Init(&h)

	for i := 0; i < n; i += 1 {
		var val int
		_, err = fmt.Scanf("%d", &val)
		if err != nil {
			return err
		}
		heap.Push(&h, val)
	}

	var k int
	_, err = fmt.Scan(&k)
	if err != nil {
		return err
	}

	res, err := h.GetKNode(k)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}
