package streamUtils

import "fmt"

func ScanInt32() (int32, error) {
	var val int32
	_, err := fmt.Scan(&val)
	return val, err
}

func ScanUInt32() (uint32, error) {
	var val uint32
	_, err := fmt.Scan(&val)
	return val, err
}

func FlushInput() {
        var flushStr string
        for nFlushed := 1 ; nFlushed != 0 ; nFlushed, _ = fmt.Scanln(&flushStr) {}
}
