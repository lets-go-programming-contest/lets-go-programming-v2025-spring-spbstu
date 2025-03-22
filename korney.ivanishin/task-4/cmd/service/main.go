package main

import (
	"errors"
	"fmt"

	"github.com/quaiion/go-practice/lru-cache/internal/CacheDir"
	"github.com/quaiion/go-practice/lru-cache/internal/Config"
	"github.com/quaiion/go-practice/lru-cache/internal/Requester"
)

var (
        errConfigFailed = errors.New("failed extracting config parameters")
)

func main() {
        configParams, err := Config.GetConfigParams()
        if err != nil {
                panic(errors.Join(errConfigFailed, err))
        }

        cacheDir := CacheDir.NewCacheDir(configParams.ReqRange,
                                         configParams.CacheCap)

        chanArr := launchRequesters(&cacheDir, configParams.NRequesters, 
                                               configParams.ReqRange,
                                               configParams.NRequests)

        hitSum := joinRequesters(chanArr, configParams.NRequesters)

        fmt.Printf("\nhit rate: %d/%d, %.2f%%\n\n", hitSum,
                   configParams.NRequests * configParams.NRequesters,
                   float64(hitSum) / float64(configParams.NRequests *
                                             configParams.NRequesters) *
                                     float64(100))
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
                requester := Requester.NewRequester(reqRange)
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
