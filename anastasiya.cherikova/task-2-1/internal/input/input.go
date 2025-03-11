// input/input.go
package input

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"task-2-1/internal/department"
)

// ReadConfig читает и валидирует N и K
func ReadConfig(scanner *bufio.Scanner) (int, error) {
	// Чтение количества отделов
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n < 1 || n > 1000 {
		return 0, fmt.Errorf("n must be 1-1000")
	}

	return n, nil
}

// ParseCondition парсит строку с температурным условием
func ParseCondition(line string) (string, int, error) {
	parts := strings.Fields(strings.TrimSpace(line))
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid condition format")
	}

	op := parts[0]
	temp, err := strconv.Atoi(parts[1])
	if err != nil || temp < 15 || temp > 30 {
		return "", 0, fmt.Errorf("temp must be 15-30")
	}

	if op != ">=" && op != "<=" {
		return "", 0, fmt.Errorf("operator must be >= or <=")
	}

	return op, temp, nil
}

// ProcessDepartments обрабатывает все отделы последовательно
func ProcessDepartments(scanner *bufio.Scanner, n int, output io.Writer) {
	for i := 0; i < n; i++ {
		// Чтение количества сотрудников
		scanner.Scan()
		k, err := strconv.Atoi(scanner.Text())
		if err != nil || k < 1 || k > 1000 {
			fmt.Println("K must be 1-1000")
			i -= 1
			continue
		}
		dept := department.New()
		processDepartment(scanner, dept, k, output)
	}
}

// Обработка одного отдела
func processDepartment(
	scanner *bufio.Scanner,
	dept *department.Department,
	k int,
	output io.Writer,
) {
	for j := 0; j < k; j++ {
		scanner.Scan()
		line := scanner.Text()

		op, temp, err := ParseCondition(line)
		if err != nil {
			fmt.Println(err)
			j -= 1
			continue
		}

		dept.Update(op, temp)
		fmt.Fprintln(output, dept.OptimalTemperature())
	}
}
