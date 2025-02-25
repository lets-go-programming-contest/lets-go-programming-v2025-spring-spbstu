package process

import (
	"errors"
	"container/heap"
	"github.com/cafe/internal/intMaxHeapIF"
	"github.com/cafe/internal/streamUtils"
)

var (
	errDataCorrupted = errors.New("internal data corrupted")
	errFailedNumConv = errors.New("number conversion error")
	errDesignPosOOR  = errors.New("designated position out of range")
)

func NewDishHeap() (*intMaxHeapIF.IntMaxHeap, error) {
	dishHeap := new(intMaxHeapIF.IntMaxHeap)

	heap.Init(dishHeap)
	return dishHeap, nil
}

func ScanDishHeap(dishHeap *intMaxHeapIF.IntMaxHeap) error {
	if dishHeap == nil {
		return errDataCorrupted
	}

	nDishes, err := streamUtils.ScanUInt32()
	if err != nil {
		return err
	}

	for i := uint32(0) ; i < nDishes ; i += 1 {
		score, err := streamUtils.ScanInt32()
		if err != nil {
			return err
		}

		heap.Push(dishHeap, score)
	}

	return nil
}

func ScanDesignPos() (uint32, error) {
	return streamUtils.ScanUInt32()
}

func GetDesignScore(dishHeap *intMaxHeapIF.IntMaxHeap, designPos uint32) (int32, error) {
	if dishHeap == nil {
		return 0, errDataCorrupted
	}
	if designPos > uint32(dishHeap.Len()) {
		return 0, errDesignPosOOR
	}

	var (
		designScore int32 = 0
		ok bool
	)
	for i := uint32(0) ; i < designPos ; i += 1 {
		designScore, ok = heap.Pop(dishHeap).(int32)
		if !ok {
			return 0, errFailedNumConv
		}
	}

	return designScore, nil
}
