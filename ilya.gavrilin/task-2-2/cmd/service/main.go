package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	
	"task-2-2/pkg/heap"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// read values
	n, nums, k, err := readInput(scanner)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// check k validity
	if k < 1 || k > n {
		fmt.Println("Некорректный порядковый номер k")
		return
	}

	// count result
	result := heap.FindKthLargest(nums, k)
	fmt.Println(result)
}

func readInput(scanner *bufio.Scanner) (int, []int, int, error) {
	// read N
	if !scanner.Scan() {
		return 0, nil, 0, scanner.Err()
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n < 1 || n > 10000 {
		return 0, nil, 0, fmt.Errorf("некорректное количество элементов")
	}

	// read array
	if !scanner.Scan() {
		return 0, nil, 0, scanner.Err()
	}
	nums, err := parseNumbers(scanner.Text(), n)
	if err != nil {
		return 0, nil, 0, err
	}

	// read k
	if !scanner.Scan() {
		return 0, nil, 0, scanner.Err()
	}
	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, 0, fmt.Errorf("некорректный формат k")
	}

	return n, nums, k, nil
}

func parseNumbers(input string, n int) ([]int, error) {
	parts := strings.Fields(input)
	if len(parts) != n {
		return nil, fmt.Errorf("ожидается %d чисел", n)
	}

	nums := make([]int, n)
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < -10000 || num > 10000 {
			return nil, fmt.Errorf("некорректное число: %s", part)
		}
		nums[i] = num
	}
	return nums, nil
}