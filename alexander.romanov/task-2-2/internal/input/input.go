package input

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type InputData struct {
	Preferences []int
	K           int
}

func ReadInput(reader io.Reader) (InputData, error) {
	scanner := bufio.NewScanner(reader)

	if !scanner.Scan() {
		return InputData{}, fmt.Errorf("failed to read number of dishes")
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return InputData{}, fmt.Errorf("error parsing number of dishes: %w", err)
	}

	if n < 1 || n > 10000 {
		return InputData{}, fmt.Errorf("number of dishes must be between 1 and 10000")
	}

	if !scanner.Scan() {
		return InputData{}, fmt.Errorf("failed to read preference values")
	}
	preferencesStr := strings.Fields(scanner.Text())
	if len(preferencesStr) != n {
		return InputData{}, fmt.Errorf("expected %d preference values, got %d", n, len(preferencesStr))
	}

	preferences := make([]int, n)
	for i, prefStr := range preferencesStr {
		pref, err := strconv.Atoi(prefStr)
		if err != nil {
			return InputData{}, fmt.Errorf("error parsing preference value: %w", err)
		}

		if pref < -10000 || pref > 10000 {
			return InputData{}, fmt.Errorf("preference value must be between -10000 and 10000")
		}

		preferences[i] = pref
	}

	if !scanner.Scan() {
		return InputData{}, fmt.Errorf("failed to read k value")
	}
	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return InputData{}, fmt.Errorf("error parsing k value: %w", err)
	}

	if k < 1 || k > n {
		return InputData{}, fmt.Errorf("k must be between 1 and %d", n)
	}

	if err := scanner.Err(); err != nil {
		return InputData{}, fmt.Errorf("error during scanning: %w", err)
	}

	return InputData{
		Preferences: preferences,
		K:           k,
	}, nil
}
