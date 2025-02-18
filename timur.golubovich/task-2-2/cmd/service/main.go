package main

import (
	"container/heap"
	"errors"
	"fmt"
	"os"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func readInt() (int, error) {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, errors.New("error in scan: " + err.Error())
	}
	return num, nil
}

func readNums() ([]int, error) {
	N, err := readInt()
	if err != nil {
		return []int{}, err
	}
	nums := []int{}
	for i := 0; i != N; i++ {
		num, err := readInt()
		if err != nil {
			return []int{}, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func solve(nums []int, k int) int {
	set := IntHeap{}
	heap.Init(&set)
	for _, num := range nums {
		heap.Push(&set, num)
	}
	for i := k; i != len(nums); i++ {
		heap.Pop(&set)
	}
	return set[0]
}

func process() error {
	nums, err := readNums()
	if err != nil {
		return err
	}
	k, err := readInt()
	if k > len(nums) || k < 0 {
		return errors.New("too big k")
	}
	if err != nil {
		return err
	}
	ans := solve(nums, k)
	fmt.Printf("%v\n", ans)
	return nil
}

func main() {
	err := process()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
	}
}
