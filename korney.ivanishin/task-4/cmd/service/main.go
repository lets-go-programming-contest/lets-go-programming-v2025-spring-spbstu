package main

import (
	"fmt"

	"github.com/quaiion/go-practice/lru-cache/internal/CacheDir"
	"github.com/quaiion/go-practice/lru-cache/internal/Config"
	"github.com/quaiion/go-practice/lru-cache/internal/Requester"
)

func main() {
        nRequesters, reqRange, cacheCap, nRequests, err := Config.GetConfigParams()
        if err != nil {
                panic(fmt.Errorf("failed extracting config parameters: %w", err))
        }

        cacheDir := CacheDir.CreateCacheDir(reqRange, cacheCap)

        chanArr := launchRequesters(&cacheDir, nRequesters, reqRange, nRequests)

        hitSum := joinRequesters(chanArr, nRequesters)

        fmt.Printf("\nhit rate: %d/%d, %.2f%%\n\n", hitSum, nRequests * nRequesters,
                   float64(hitSum) / float64(nRequests * nRequesters) * float64(100))
}

func goroutRequestN(requester Requester.Requester, cd *CacheDir.CacheDir, n uint32,  ch chan<- uint32) {
        nHits, err := requester.RequestN(cd, n)
        if err != nil {
                panic(err)
        }

        ch <- nHits
}

func launchRequesters(cd *CacheDir.CacheDir, nRequesters uint32, reqRange uint32, nRequests uint32) *[]chan uint32 {
        chanArr := make([]chan uint32, 0, nRequesters)

        for i := uint32(0) ; i < nRequesters ; i += 1 {
                requester := Requester.CreateRequester(reqRange)
                chanArr = append(chanArr, make(chan uint32))
                go goroutRequestN(requester, cd, nRequests, chanArr[i])
        }

        return &chanArr
}

func joinRequesters(chanArr *[]chan uint32, nRequesters uint32) uint32 {
        var hitSum uint32 = 0

        for i := uint32(0) ; i < nRequesters ; i += 1 {
                hitSum += <-((*chanArr)[i])
        }

        return hitSum
}
