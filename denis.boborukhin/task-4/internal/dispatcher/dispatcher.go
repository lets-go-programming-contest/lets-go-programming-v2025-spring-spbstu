package dispatcher

import (
	"sync"

	"github.com/denisboborukhin/downloader/internal/downloader"
)

type Dispatcher struct {
	dl         *downloader.Downloader
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		dl:         downloader.NewDownloader(),
		maxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Start(files []downloader.File) {
	semaphore := make(chan struct{}, d.maxWorkers)
	var wg sync.WaitGroup

	for _, file := range files {
		semaphore <- struct{}{}
		wg.Add(1)

		go func(f downloader.File) {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			d.dl.Download(f)
		}(file)
	}

	close(semaphore)
	wg.Wait()
}
