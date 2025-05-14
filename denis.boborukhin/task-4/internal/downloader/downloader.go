package downloader

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type File struct {
	URL  string
	Size int
}

type Downloader struct {
	mu     sync.Mutex
	active map[string]bool
}

func NewDownloader() *Downloader {
	return &Downloader{
		active: make(map[string]bool),
	}
}

func (d *Downloader) Download(file File) error {
	d.mu.Lock()
	if d.active[file.URL] {
		d.mu.Unlock()
		return fmt.Errorf("file %s is already downloading", file.URL)
	}
	d.active[file.URL] = true
	d.mu.Unlock()

	time.Sleep(time.Duration(rand.Intn(500)+file.Size/10) * time.Millisecond)
	fmt.Printf("file %s downloaded (%d KB)\n", file.URL, file.Size)

	d.mu.Lock()
	delete(d.active, file.URL)
	d.mu.Unlock()
	return nil
}
