// main.go
package main

import (
	"task-2-2/internal/io"
	"task-2-2/internal/logic"
)

func main() {
	nums, k := io.ReadInput()
	result := logic.FindKthLargest(nums, k)
	io.PrintResult(result)
}
