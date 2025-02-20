package sliceReader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Read(N uint) ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите количество баллов предпочтения блюд:")
	scanner.Scan()
	input := scanner.Text()

	numbers := strings.Fields(input)

	if len(numbers) != int(N) {
		fmt.Printf("Ошибка: ожидается %d чисел, но введено %d\n", N, len(numbers))
		err := errors.New("too many args")
		return []int{}, err
	}

	intNumbers := make([]int, 0, N)

	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			err := errors.New("arg not int")
			return []int{}, err
		}

		if num < -10000 || num > 10000 {
			err := errors.New("bad arg")
			return []int{}, err
		}
		intNumbers = append(intNumbers, num)
	}

	return intNumbers, nil
}
