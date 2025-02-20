package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/kseniadobrovolskaia/task-2-1/internal/interval"
)

const (
	lowBound = 15
	upBound  = 30
)

func main() {
	result, err := runPlaceInSun()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nРезультат: %v\n", result)
}

func runPlaceInSun() ([]int, error) {
	fmt.Print("Введите количество отделов: ")
	n, err := readNumber()
	if err != nil {
		return nil, err
	}
	result := make([]int, 0)

	for i := 0; i < n; i++ {
		fmt.Printf("\nВведите количество сотрудников в отделе %d: ", i+1)
		m, err := readNumber()
		if err != nil {
			return nil, err
		}
		optT, err := getOptimalTempForDepart(m)
		if err != nil {
			return nil, err
		}
		result = append(result, optT...)

	}
	return result, nil
}

func getOptimalTempForDepart(people int) ([]int, error) {
	result := make([]int, 0, people)
	optimTemp := interval.NewIntervalValue(lowBound, upBound)
	for human := 0; human < people; human++ {
		fmt.Printf("Введите температурную границу сотрудника %d: ", human+1)
		var lessOrBigger string
		var t int
		fmt.Scanf("%s %d", &lessOrBigger, &t)
		if lessOrBigger != "<=" && lessOrBigger != ">=" {
			return nil, errors.New("Температурная граница должна начинаться со знаков \"<=\" или \">=\", после которых через пробел следует числовое значение.\n")
		}
		switch lessOrBigger {
		case "<=":
			optimTemp.LessThan(t)
		case ">=":
			optimTemp.BiggerThan(t)
		default:
			return nil, errors.New("Не можем здесь оказаться")
		}
		result = append(result, optimTemp.Value)
	}
	return result, nil
}

func readNumber() (num int, err error) {
	_, err = fmt.Scan(&num)
	if err != nil {
		err = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	return
}
