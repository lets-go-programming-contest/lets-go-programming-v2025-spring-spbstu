package CacheDir

import (
	"container/list"
	"fmt"
	"sync"
)

type CacheDir struct {
        list *list.List
        dir  []*list.Element
        rang uint32
        size uint32
        cap  uint32
        mtx  sync.Mutex
}

func NewCacheDir(reqRange uint32, cacheCap uint32) CacheDir {
        return CacheDir{
                list: list.New(),
                dir: make([]*list.Element, int(reqRange)),
                rang: reqRange,
                size: 0,
                cap: cacheCap,
                mtx: sync.Mutex{},
        }
}

func (cd *CacheDir) GetRequest(id uint32) (bool, error) {
        if id >= cd.rang {
                return false, fmt.Errorf("request id %d out of predefined range of %d",
                                         id, cd.rang)
        }

        cd.mtx.Lock()
        if cd.dir[id] != nil {
                err := cd.processHit(id)
                cd.mtx.Unlock()
                if err != nil {
                        return true, fmt.Errorf("failed processing a hit on request id %d: %w",
                                                id, err)
                }

                return true, nil
        } else {
                err := cd.processMiss(id)
                cd.mtx.Unlock()
                if err != nil {
                        return false, fmt.Errorf("failed processing a miss on request id %d: %w",
                                                 id, err)
                }

                return false, nil
        }
}

func (cd *CacheDir) processHit(id uint32) error {
        cd.list.MoveToFront(cd.dir[id])

        return nil
}

func (cd *CacheDir) processMiss(id uint32) error {
        cd.list.PushFront(id)
        cd.dir[id] = cd.list.Front()

        if cd.size < cd.cap {
                cd.size += 1
        } else {
                removed, ok := cd.list.Remove(cd.list.Back()).(uint32)
                if !ok {
                        panic("data in cachedir list corrupted")
                }

                cd.dir[removed] = nil
        }

        return nil
}
