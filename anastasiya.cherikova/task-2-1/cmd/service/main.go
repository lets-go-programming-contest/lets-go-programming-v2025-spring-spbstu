// main.go
package main

import (
	"bufio"
	"fmt"
	"os"

	"task-2-1/internal/input"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Чтение конфигурации
	n, err := input.ReadConfig(scanner)
	if err != nil {
		fmt.Println(-1)
		return
	}

	// Запуск обработки данных
	input.ProcessDepartments(scanner, n, os.Stdout)
}
