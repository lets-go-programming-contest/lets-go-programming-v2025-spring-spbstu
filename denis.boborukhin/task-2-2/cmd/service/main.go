package main

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []int

func (ch PriorityQueue) Len() int {
	return len(ch)
}

func (ch PriorityQueue) Less(i, j int) bool {
	return ch[i] < ch[j]
}

func (ch PriorityQueue) Swap(i, j int) {
	ch[i], ch[j] = ch[j], ch[i]
}

func (ch *PriorityQueue) Push(x interface{}) {
	if val, ok := x.(int); ok {
		*ch = append(*ch, val)
	} else {
		fmt.Println("Error push: invalid type")
	}
}

func (ch *PriorityQueue) Pop() interface{} {
	old := *ch
	n := len(old)
	x := old[n-1]
	*ch = old[:n-1]
	return x
}

func scanInteger() (int, error) {
	var value int
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, fmt.Errorf("scan error: %v", err)
	}
	return value, nil
}

func scanIntegers() ([]int, error) {
	count, err := scanInteger()
	if err != nil {
		return nil, err
	}
	values := make([]int, 0, count)
	for i := 0; i < count; i++ {
		value, err := scanInteger()
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func findKthBiggest(values []int, k int) int {
	h := PriorityQueue(values)
	heap.Init(&h)
	for i := k; i < len(values); i++ {
		heap.Pop(&h)
	}
	return h[0]
}

func execute() error {
	values, err := scanIntegers()
	if err != nil {
		return err
	}
	k, err := scanInteger()
	if err != nil {
		return err
	}
	if k < 0 || k > len(values) {
		return fmt.Errorf("invalid k value")
	}
	result := findKthBiggest(values, k)
	fmt.Println(result)
	return nil
}

func main() {
	if err := execute(); err != nil {
		fmt.Println("Error occurred:", err)
	}
}
