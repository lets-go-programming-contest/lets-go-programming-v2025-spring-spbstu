package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/realFrogboy/task-4/internal/lockedrange"
)

func doSync(queue *lockedrange.LockingRangeAdapter) {
	for idx := range 100 {
		queue.Add(idx)
	}
}

func doAsync(queue *[]int) {
	for idx := range 100 {
		*queue = append(*queue, idx)
	}
}

func main() {
	isSync := flag.Bool("sync", false, "")
	flag.Parse()

	var data []int

	if *isSync {
		queue := lockedrange.NewLockingRangeAdapter(data)

		var wg sync.WaitGroup
		for _ = range 10 {
			wg.Add(1)

			go func() {
				defer wg.Done()
				doSync(queue)
			}()
		}

		wg.Wait()

		fmt.Println(queue.Len())
	} else {
		var wg sync.WaitGroup
		for _ = range 10 {
			wg.Add(1)

			go func() {
				defer wg.Done()
				doAsync(&data)
			}()
		}

		wg.Wait()

		fmt.Println(len(data))
	}
}
