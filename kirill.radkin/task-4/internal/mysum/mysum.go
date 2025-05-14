package mysum

import "github.com/yanelox/task-4/internal/syncsum"


func countSumPart (N_start int, N_end int, sum *syncsum.SyncSum, channel chan bool) {
	var partSum float64 = 0
	for i := N_start; i <= N_end; i++ {
		partSum += 1 / float64(i)
	}

	sum.Increase(partSum)

	channel <- true
}

func CountSum (N int, goroutines_count int) float64 {
	channel := make(chan bool)
	asyncSum := syncsum.NewSyncSum()

	var default_count int = N / goroutines_count
	var last_count int = N - default_count * (goroutines_count - 1)

	if goroutines_count == 1 {
		go countSumPart(1, N, asyncSum, channel)

		<- channel
	} else {
		for i := 0; i < goroutines_count - 1; i++ {
			go countSumPart(1 + i * default_count, 1 + (i + 1) * default_count - 1, asyncSum, channel);
		}

		go countSumPart(N + 1 - last_count, N, asyncSum, channel);

		for i := 0; i < goroutines_count; i++ {
			<- channel
		}
	}

	return asyncSum.Get()
}