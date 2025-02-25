package choice

import (
	"container/heap"
	"fmt"
	"log"
	IntHeap "task-2-2/cmd/intHeap"
)

func Run() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		handleError(err)
	}

	h := make(IntHeap.IntHeap, n)
	heap.Init(&h)

	for i := 0; i < n; i += 1 {
		var val int
		_, err = fmt.Scanf("%d", &val)
		if err != nil {
			handleError(err)
		}
		heap.Push(&h, val)
	}

	var k int
	_, err = fmt.Scan(&k)
	if err != nil {
		handleError(err)
	}

	res, err := h.GetKNode(k)
	if err != nil {
		handleError(err)
	}

	fmt.Println(res)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
