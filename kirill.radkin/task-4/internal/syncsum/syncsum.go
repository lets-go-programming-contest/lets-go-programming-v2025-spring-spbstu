package syncsum

import "sync"

type SyncSum struct {
	sum float64
	mutex *sync.Mutex
}

func NewSyncSum () (*SyncSum) {
	var asyncSum SyncSum

	var mutex sync.Mutex
	asyncSum.mutex = &mutex

	return &asyncSum
}

func (syncSum *SyncSum) Increase (num float64) {
	syncSum.mutex.Lock()
	syncSum.sum += num
	syncSum.mutex.Unlock()
}

func (syncSum *SyncSum) Get () (float64) {
	return syncSum.sum
}