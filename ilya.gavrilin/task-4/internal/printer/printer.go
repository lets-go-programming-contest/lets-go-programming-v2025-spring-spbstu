package printer

import "fmt"

func LogRequest(url, data string, fromCache bool) {
	status := "CACHE MISS"
	if fromCache {
		status = "CACHE HIT"
	}
	fmt.Printf("[%s] %s -> %s\n", status, url, data)
}