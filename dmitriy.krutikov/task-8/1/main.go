package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	data := []int{5, 1, 5, 6666, 22, 11, 90}
	fmt.Println("Original array:", data)

	bubbleSort(data)

	fmt.Println("Sorted array:  ", data)
}