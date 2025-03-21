package lockedrange

import (
	"sync"
)

type LockingRangeAdapter struct {
	m_data  []int
	m_mutex sync.Mutex
}

func NewLockingRangeAdapter(data []int) *LockingRangeAdapter {
	adapter := new(LockingRangeAdapter)
	adapter.m_data = data
	return adapter
}

func (adapter *LockingRangeAdapter) Add(value int) {
	adapter.m_mutex.Lock()
	defer adapter.m_mutex.Unlock()

	adapter.m_data = append(adapter.m_data, value)
}

func (adapter *LockingRangeAdapter) Len() int {
	return len(adapter.m_data)
}
