package main

import (
	"container/heap"
	"errors"
	"fmt"
	"log"

	"github.com/kseniadobrovolskaia/task-2-2/internal/intheap"
)

func main() {
	result, err := runProblemOfChoice()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nРезультат: %v\n", result)
}

func runProblemOfChoice() (int, error) {
	fmt.Print("Введите количество блюд (1 - 10000): ")
	countDish, err := readNumber()
	if err != nil {
		return 0, err
	}
	rating := make(intheap.IntHeap, 0, countDish)
	prating := &rating
	heap.Init(prating)
	fmt.Printf("\nВведите через пробел рeйтинги блюд (-10000, 10000): ")

	for i := 0; i < countDish; i++ {
		rat, err := readNumber()
		if err != nil {
			return 0, err
		}
		heap.Push(prating, rat)

	}
	fmt.Print("Введите порядковый номер k-го по предпочтению блюда: ")
	k, err := readNumber()
	if err != nil {
		return 0, err
	}
	if k < 0 || k > countDish {
		return 0, errors.New("Порядковый номер должен быть положительный и меньше или равен числу блюд")
	}
	for ; k < countDish; k++ {
		heap.Pop(prating)
	}
	return rating[0], nil
}

func readNumber() (num int, err error) {
	_, err = fmt.Scan(&num)
	if err != nil {
		err = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	return
}
