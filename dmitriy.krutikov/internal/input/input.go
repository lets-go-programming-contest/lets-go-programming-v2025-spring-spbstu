package input

import (
    "errors"
    "fmt"
)

func InputNumber() (int, error) {
    var number int
    _, err := fmt.Scanln(&number)

    if err != nil {
        return 0, errors.New("Incorrect number\n")
    }

    return number, nil
}