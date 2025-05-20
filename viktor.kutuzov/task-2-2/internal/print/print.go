package print

import "fmt"

func PrintAns(val1, val2, ans float64, op string) error {
	_, err := fmt.Println("Result:", val1, op, val2, "=", ans)
	return err
}
