package sieve

import "errors"

var (
        errIndexOutOfBound   = errors.New("critical! sieve somehow indexed out of bounds")
        errSieveUpdateFailed = errors.New("sieve update failed")
)

func FindPrimes(edge uint32) ([]uint32, error) {
        // false => (potentially) prime, true => non-prime
        sieve := make([]bool, edge)
        
        primes := make([]uint32, 0)
        
        for num := uint32(2) ; num <= edge ; num += 1 {
                if !sieve[num - 2] {
                        primes = append(primes, num)

                        err := updateSieve(sieve, num)
                        if err != nil {
                                return nil, errors.Join(errSieveUpdateFailed, err)
                        }
                }
        }

        return primes, nil
}

func updateSieve(sieve []bool, newPrime uint32) error {
        if newPrime > uint32(len(sieve)) {
                return errIndexOutOfBound
        }

        for num := newPrime ; num <= uint32(len(sieve)) + 1 ; num += newPrime {
                sieve[num - 2] = true
        }

        return nil
}
