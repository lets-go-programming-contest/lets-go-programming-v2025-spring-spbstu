package loader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// LoadData loads data from a local file or a remote URL.
func LoadData(path string) (io.ReadCloser, error) {
	if strings.HasPrefix(path, "https://") {
		return fetchFromURL(path)
	}
	return loadFromFile(path)
}

// fetchFromURL downloads XML data from a given URL.
func fetchFromURL(url string) (io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers to avoid 403 errors
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from URL %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status: %d", resp.StatusCode)
	}

	return resp.Body, nil
}

// loadFromFile reads XML data from a local file.
func loadFromFile(path string) (io.ReadCloser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}
	return file, nil
}
