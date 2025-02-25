package main

import (
	"fmt"
	"github.com/cafe/internal/process"
	"github.com/cafe/internal/streamUtils"
)

func main() {
	dishHeap, err := process.NewDishHeap()
	if err != nil {
		streamUtils.FlushInput()
		fmt.Println(err)
		return
	}

	err = process.ScanDishHeap(dishHeap)
	if err != nil {
		streamUtils.FlushInput()
		fmt.Println(err)
		return
	}

	designPos, err := process.ScanDesignPos()
	if err != nil {
		streamUtils.FlushInput()
		fmt.Println(err)
		return
	}

	designScore, err := process.GetDesignScore(dishHeap, designPos)
	if err != nil {
		streamUtils.FlushInput()
		fmt.Println(err)
		return
	}

	fmt.Println(designScore)
}
