package lockedrange

import (
	"sync"
)

type LockingRangeAdapter struct {
	data  []int
	mutex sync.Mutex
}

func NewLockingRangeAdapter(data []int) *LockingRangeAdapter {
	adapter := new(LockingRangeAdapter)
	adapter.data = data
	return adapter
}

func (adapter *LockingRangeAdapter) Add(value int) {
	adapter.mutex.Lock()
	defer adapter.mutex.Unlock()

	adapter.data = append(adapter.data, value)
}

func (adapter *LockingRangeAdapter) Len() int {
	adapter.mutex.Lock()
	defer adapter.mutex.Unlock()

	return len(adapter.data)
}
