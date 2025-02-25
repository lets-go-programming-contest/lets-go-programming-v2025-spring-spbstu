package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	minTemp = 15
	maxTemp = 30
)

type Department struct {
	min int
	max int
}

func readIntInRange(scanner *bufio.Scanner, min, max int) (int, bool) {
	if !scanner.Scan() {
		return 0, false
	}
	value, err := strconv.Atoi(scanner.Text())
	if err != nil || value < min || value > max {
		return 0, false
	}
	return value, true
}

func readTemperatureConstraint(scanner *bufio.Scanner) (string, int, bool) {
	if !scanner.Scan() {
		return "", 0, false
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) != 2 {
		return "", 0, false
	}

	sign, tempStr := parts[0], parts[1]
	temp, err := strconv.Atoi(tempStr)
	if err != nil || temp < minTemp || temp > maxTemp || (sign != ">=" && sign != "<=") {
		return "", 0, false
	}

	return sign, temp, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	N, valid := readIntInRange(scanner, 1, 1000)
	if !valid {
		log.Fatal("Invalid number of departments")
	}

	for i := 0; i < N; i++ {
		K, valid := readIntInRange(scanner, 1, 1000)
		if !valid {
			log.Fatal("Invalid number of employees in department ", i+1)
		}

		department := Department{minTemp, maxTemp}

		for j := 0; j < K; j++ {
			sign, temp, valid := readTemperatureConstraint(scanner)
			if !valid {
				log.Fatal("Invalid temperature constraint")
			}

			if sign == "<=" {
				if temp < department.max {
					department.max = temp
				}
			} else if sign == ">=" {
				if temp > department.min {
					department.min = temp
				}
			}

			if department.min <= department.max {
				fmt.Println(department.min)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
