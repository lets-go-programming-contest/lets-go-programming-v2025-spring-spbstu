package main

import (
	"errors"

	"github.com/quaiion/go-practice/sieve/internal/sieve"
	"github.com/quaiion/go-practice/sieve/internal/streamutils"
)

var (
        errScanFailed  = errors.New("input scan failed")
        errSieveFailed = errors.New("failed to apply the sieve")
        errPrintFailed = errors.New("failed to print the primes")
)

func main() {
        edge, err := streamutils.ScanUInt32()
        if err != nil {
                streamutils.FlushInput()
                panic(errors.Join(errScanFailed, err))
        }

        primes, err := sieve.FindPrimes(edge)
        if err != nil {
                panic(errors.Join(errSieveFailed, err))
        }

        err = streamutils.PrintSlice(primes)
        if err != nil {
                panic(errors.Join(errPrintFailed, err))
        }
}
