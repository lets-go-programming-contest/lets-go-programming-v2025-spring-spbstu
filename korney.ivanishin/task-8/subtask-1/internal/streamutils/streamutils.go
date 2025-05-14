package streamutils

import (
	"errors"
	"fmt"
)

var errSlicePrintFailed = errors.New("failed printing a slice")

func ScanUInt32() (uint32, error) {
	var v uint32
	_, err := fmt.Scan(&v)
	return v, err
}

func PrintSlice[T any](s []T) error {
        for _, v := range s {
                _, err := fmt.Printf("%v ", v)
                if err != nil {
                        return errors.Join(errSlicePrintFailed, err)
                }
        }

        fmt.Println(``)
        return nil
}

func FlushInput() {
        var flushStr string
        for nFlushed := 1 ; nFlushed != 0 ; nFlushed, _ = fmt.Scanln(&flushStr) {}
}
