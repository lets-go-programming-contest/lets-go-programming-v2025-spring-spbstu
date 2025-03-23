package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/quaiion/go-practice/lru-cache/internal/CacheDir"
	"github.com/quaiion/go-practice/lru-cache/internal/Config"
	"github.com/quaiion/go-practice/lru-cache/internal/Requester"
	"golang.org/x/sync/errgroup"
)

var (
        errConfigFailed = errors.New("failed extracting config parameters")
        errRequestSimFailed = errors.New("failed simulating requests")
)

func main() {
        configParams, err := Config.GetConfigParams()
        if err != nil {
                panic(errors.Join(errConfigFailed, err))
        }

        cacheDir := CacheDir.NewCacheDir(configParams.ReqRange,
                                         configParams.CacheCap)

        chanArr, group := launchRequesters(&cacheDir,
                                           configParams.NRequesters, 
                                           configParams.ReqRange,
                                           configParams.NRequests)

        hitSum, err := joinRequesters(chanArr, configParams.NRequesters, group)
        if err != nil {
                panic(errors.Join(errRequestSimFailed, err))
        }

        fmt.Printf("\nhit rate: %d/%d, %.2f%%\n\n", hitSum,
                   configParams.NRequests * configParams.NRequesters,
                   float64(hitSum) / float64(configParams.NRequests *
                                             configParams.NRequesters) *
                                     float64(100))
}

func goroutRequestN(requester Requester.Requester, cd *CacheDir.CacheDir, n uint32, ch chan<- uint32) error {
        nHits, err := requester.RequestN(cd, n)
        if err != nil {
                // without this the system deadlocks waiting
                // for the channel in an error scenario
                ch <- 0

                return err
        }

        ch <- nHits
        return nil
}

func launchRequesters(cd *CacheDir.CacheDir, nRequesters uint32, reqRange uint32, nRequests uint32) (*[]chan uint32, *errgroup.Group) {
        chanArr := make([]chan uint32, 0, nRequesters)

        group, _ := errgroup.WithContext(context.Background())

        for i := uint32(0) ; i < nRequesters ; i += 1 {
                requester := Requester.NewRequester(reqRange)
                chanArr = append(chanArr, make(chan uint32))
                group.Go(func() error {
                        return goroutRequestN(requester, cd, nRequests, chanArr[i])
                })
        }

        return &chanArr, group
}

var errRequestFailed = errors.New("failed creating or processing a request")

func joinRequesters(chanArr *[]chan uint32, nRequesters uint32, group *errgroup.Group) (uint32, error) {
        // postponed summation not required as zero is
        // returned from goroutine in an error case
        var hitSum uint32 = 0
        for i := uint32(0) ; i < nRequesters ; i += 1 {
                hitSum += <-((*chanArr)[i])
        }
        
        err := group.Wait()
        if err != nil {
                return 0, errors.Join(errRequestFailed, err)
        }

        return hitSum, nil
}
