package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"task-2-1/pkg/temperature"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// read number of departments
	n := readValidInt(scanner, 1, 1000, "Неверное количество отделов")

	for i := 0; i < n; i++ {
		processDepartment(scanner)
	}
}

func processDepartment(scanner *bufio.Scanner) {
	tr := temperature.NewTemperatureRange()

	// read number of employee
	k := readValidInt(scanner, 1, 1000, "Неверное количество сотрудников")

	for j := 0; j < k; j++ {
		processEmployee(scanner, tr)
	}
}

func processEmployee(scanner *bufio.Scanner, tr *temperature.TemperatureRange) {
	constraint, value, ok := readConstraint(scanner)
	if !ok {
		fmt.Println(-1)
		return
	}

	// applying constraits
	result := tr.ApplyConstraint(constraint, value)
	fmt.Println(result)
}

func readValidInt(scanner *bufio.Scanner, min, max int, errMsg string) int {
	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(input)
		
		if err == nil && num >= min && num <= max {
			return num
		}
		fmt.Printf("%s (допустимый диапазон %d-%d)\n", errMsg, min, max)
	}
}

// read temperature constraits with error handling
func readConstraint(scanner *bufio.Scanner) (string, int, bool) {
	for {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		
		if len(parts) != 2 {
			fmt.Println("Ошибка формата: ожидается [операция] [значение]")
			continue
		}

		op := parts[0]
		value, err := strconv.Atoi(parts[1])
		
		if (op != ">=" && op != "<=") || err != nil || value < 15 || value > 30 {
			fmt.Println("Некорректное ограничение. Формат: >=/<= [15-30]")
			continue
		}

		return op, value, true
	}
}