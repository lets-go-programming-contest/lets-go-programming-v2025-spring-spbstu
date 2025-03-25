package dispatcher

import (
	"github.com/denisboborukhin/downloader/internal/downloader"
)

type Dispatcher struct {
	dl         *downloader.Downloader
	semaphore  chan struct{}
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		dl:         downloader.NewDownloader(),
		maxWorkers: maxWorkers,
		semaphore:  make(chan struct{}, maxWorkers),
	}
}

func (d *Dispatcher) Start(files []downloader.File) {
	for _, file := range files {
		d.semaphore <- struct{}{}

		go func(f downloader.File) {
			defer func() {
				<-d.semaphore
			}()

			d.dl.Download(f)
		}(file)
	}

	for i := 0; i < cap(d.semaphore); i++ {
		d.semaphore <- struct{}{}
	}
	close(d.semaphore)
}
