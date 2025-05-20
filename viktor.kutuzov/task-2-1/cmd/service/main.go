package main

import (
	"fmt"
	"log"

	"github.com/kutuzov.viktor/task-2-1/internal/calculate"
	"github.com/kutuzov.viktor/task-2-1/internal/uintReader"
)

func main() {
	fmt.Print("Введите количество отделов: ")
	var N uint64
	N, err := uintReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	if N == 0 {
		log.Fatal("Ноль отделов -> задача отсутвует")
	}

	for i := uint64(0); i < N; i++ {
		fmt.Printf("Количество сотрудников в отделе %d: ", i)
		k, err := uintReader.Read()
		if err != nil {
			log.Fatal(err)
		}

		if k == 0 {
			log.Fatal("Не существует отдела без сотрудников")
		}

		temp, err := calculate.CalcSuitableTemp(i, k)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Оптимальная температура в отделе", i, ":", temp, "°C")

	}

}
