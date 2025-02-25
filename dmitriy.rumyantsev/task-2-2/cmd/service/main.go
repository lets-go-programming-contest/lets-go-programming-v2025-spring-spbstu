package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dmitriy.rumyantsev/task-2-2/pkg/min_heap"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, err := strconv.Atoi(scanner.Text())
	if err != nil || N < 1 || N > 10000 {
		fmt.Println("Invalid number of dishes")
		log.Fatal()
	}

	scanner.Scan()
	numbers := strings.Fields(scanner.Text())
	if len(numbers) != N {
		fmt.Println("Number of ratings does not match number of dishes")
		log.Fatal()
	}

	ratings := make([]int, N)
	for i := 0; i < N; i++ {
		value, err := strconv.Atoi(numbers[i])
		if err != nil || value < -10000 || value > 10000 {
			fmt.Println("Invalid dish rating")
			log.Fatal()
		}
		ratings[i] = value
	}

	scanner.Scan()
	K, err := strconv.Atoi(scanner.Text())
	if err != nil || K < 1 || K > N {
		fmt.Println("Invalid k value")
		log.Fatal()
	}

	fmt.Println(min_heap.FindKthLargest(ratings, K))
}
