package main

import (
	"fmt"
	"net/http"
	"sync"

	"task-4/internal/cache"
	"task-4/internal/printer"
	"task-4/internal/reader"
)

const useSync = true // Enable/Disable sync

func main() {
	c := cache.NewCache(useSync)
	var wg sync.WaitGroup

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		go func() {
			defer wg.Done()

			url := reader.GenerateRandomURL()

			// check if data cached
			if data, found := c.Get(url); found {
				printer.LogRequest(url, data, true)
				fmt.Fprintf(w, "Cached Response: %s\n", data)
				return
			}

			// if not in cache: emulate "request" and set
			data := cache.SimulateExternalFetch(url)
			c.Set(url, data)
			printer.LogRequest(url, data, false)
			fmt.Fprintf(w, "Fetched Response: %s\n", data)
		}()
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
	wg.Wait()
}