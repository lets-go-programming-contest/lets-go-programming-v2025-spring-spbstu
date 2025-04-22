package main

import (
	"fmt"
	"net/http"

	"task-4/internal/cache"
	"task-4/internal/printer"
	"task-4/internal/reader"
)

const useSync = false // Enable/Disable sync

func main() {
	c := cache.NewCache(useSync)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := reader.GenerateRandomURL()

		// Check if data is cached
		if data, found := c.Get(url); found {
			printer.LogRequest(url, data, true)
			fmt.Fprintf(w, "Cached Response: %s\n", data)
			return
		}

		// If not in cache: fetch and store
		data := cache.SimulateExternalFetch(url)
		c.Set(url, data)
		printer.LogRequest(url, data, false)
		fmt.Fprintf(w, "Fetched Response: %s\n", data)
	})

	fmt.Println("Starting server on :8080")
	//it always returns non-nil errror, so leave it unchecked
	// Used ionly as primitive demo, not use in real project without CLose
	http.ListenAndServe(":8080", nil)
}
