// io/io.go
package io

import "fmt"

// ReadInput читает входные данные
func ReadInput() ([]int, int) {
	var n, k int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Scan(&k)
	if k < 1 || k > n {
		panic("Некорректный порядковый номер k")
	}
	return nums, k
}

// PrintResult выводит результат
func PrintResult(result int) {
	fmt.Println(result)
}
