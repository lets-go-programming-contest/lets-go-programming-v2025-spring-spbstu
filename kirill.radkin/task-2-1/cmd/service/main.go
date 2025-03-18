package main

import "fmt"

func main() {
	var res []int

	var numDepartments int
	_, err := fmt.Scan(&numDepartments)

	if err != nil {
		panic("Incorrect input")
	}

	for i := 0; i < numDepartments; i++ {
		tempMin := 15
		tempMax := 30
		fail := false

		var numEmployee int
		_, err := fmt.Scan(&numEmployee)

		if err != nil {
			panic("Incorrect input")
		}

		for j := 0; j < numEmployee; j++ {
			var str string
			var temp int
			_, err = fmt.Scan(&str, &temp)

			if err != nil || (str != ">=" && str != "<=") {
				panic("Incorrect input")
			}

			if str == ">=" {
				if temp > tempMax {
					fail = true
				}

				if temp > tempMin {
					tempMin = temp
				}
			}

			if str == "<=" {
				if temp < tempMax {
					tempMax = temp
				}

				if temp < tempMin {
					fail = true
				}
			}

			if tempMin > tempMax {
				fail = true
			}

			if fail {
				res = append(res, -1)
			} else {
				res = append(res, tempMin)
			}
		}
	}

	for i := range res {
		fmt.Println(res[i])
	}
}
