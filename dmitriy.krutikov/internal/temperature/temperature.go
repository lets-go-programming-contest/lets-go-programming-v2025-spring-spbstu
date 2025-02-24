package temperature

import (
	"bufio"
    "errors"
    "fmt"
	"os"
    "office-temperature/internal/input"
)

const (
    MinTemp = 15
    MaxTemp = 30
)

func caclulatePersonTemperature(minTemp, maxTemp int, constraint string) (int, int, error) {
	var op string
    var value int

	_, err := fmt.Sscanf(constraint, "%s %d", &op, &value)

	if err != nil {
        return minTemp, maxTemp, errors.New("invalid constraint format")
    }

	if value > MaxTemp || value < MinTemp {
		return minTemp, maxTemp, errors.New("temperature out of range")
	}

	switch op {
    case "<=":
        if value < maxTemp {
            maxTemp = value
        }
    case ">=":
        if value > minTemp {
            minTemp = value
        }
    default:
        return minTemp, maxTemp, errors.New("unsupported operation")
    }

    return minTemp, maxTemp, nil

}

func caclulateDepartamentTemperature(K int) (error) {
	reader := bufio.NewReader(os.Stdin)

	newMin := MinTemp
    newMax := MaxTemp

	for i := 0; i < K; i++ {
        fmt.Printf("Enter constraint for employee %d (for example, '>= 20'): ", i+1)
        constraint, err := reader.ReadString('\n')

        if err != nil {
            return errors.New("Incorrect number\n")
        }

		newMin, newMax, err = caclulatePersonTemperature(newMin, newMax, constraint)

        if err != nil {
            return err
        }

		if newMin > newMax {
            fmt.Println("Optimal temperature: -1\n")
        } else {
            fmt.Printf("Optimal temperature: %d\n", newMin)
        }
	}

	return nil
}




func Run() error {
	fmt.Print("Enter the number of departments: ");
	N, err := input.InputNumber()
    if err != nil {
        return err
    }
	
	for i := 0; i < N; i++ {
		fmt.Print("Enter the number of employees: ");
		K, err := input.InputNumber()

		if err != nil {
			return err
		}

		err = caclulateDepartamentTemperature(K)

		if err != nil {
			return err
		}
	}


	return nil


}