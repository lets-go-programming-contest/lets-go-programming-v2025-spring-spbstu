package stream_utils

import "fmt"

func FlushInput() {
        var flushStr string
        for nFlushed := 1; nFlushed != 0; nFlushed, _ = fmt.Scanln(&flushStr) {}
}
