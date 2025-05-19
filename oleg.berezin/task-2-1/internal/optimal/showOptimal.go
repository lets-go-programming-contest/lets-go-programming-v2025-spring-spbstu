package optimal

import (
	"fmt"
)

func ShowOptimal(interval OptInt) {
	_, err := fmt.Printf("%d\n", interval.T1)
	if err != nil {
		fmt.Printf("error during printf\n")
		panic(err)
	}
}
