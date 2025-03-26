package main

import (
	"fmt"

	"github.com/denisboborukhin/downloader/internal/dispatcher"
	"github.com/denisboborukhin/downloader/internal/downloader"
)

func main() {
	files := []downloader.File{
		{URL: "http://example.com/file1.zip", Size: 1024},
		{URL: "http://example.com/file2.zip", Size: 2048},
		{URL: "http://example.com/file3.zip", Size: 512},
		{URL: "http://example.com/file4.mp3", Size: 768},
		{URL: "http://example.com/file5.pdf", Size: 256},
	}

	fmt.Println("Start downloading...")

	const numWorkers = 3
	dp := dispatcher.NewDispatcher(numWorkers)
	dp.Start(files)

	fmt.Println("All done!")
}
