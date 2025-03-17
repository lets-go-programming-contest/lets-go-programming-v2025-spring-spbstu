package calculator

import (
	"fmt"
	"math/big"
	"sync"
)

func factorial(count uint64) *big.Int {
	fact := big.NewInt(1)
	start := uint64(1)
	end := count + 1
	for i := start; i != end; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	return fact
}

func computeSegment(
	start uint64,
	end uint64,
	wg *sync.WaitGroup,
	result chan<- *big.Rat) {
	defer wg.Done()
	sum := big.NewRat(0, 1)
	tmp := big.NewRat(0, 1)
	for i := start; i != end; i++ {
		fact := factorial(i)
		tmp.SetFrac(big.NewInt(1), fact)
		sum.Add(sum, tmp)
	}
	result <- sum
}

func delegate(itersCount uint64, nodesCount uint64) *big.Rat {
	var wg sync.WaitGroup
	result := make(chan *big.Rat, nodesCount)

	for rank := uint64(0); rank != nodesCount; rank++ {
		start := itersCount / nodesCount * rank
		end := itersCount / nodesCount * (rank + 1)
		if rank+1 == nodesCount {
			end += itersCount % nodesCount
		}
		wg.Add(1)
		go computeSegment(start, end, &wg, result)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	exp := big.NewRat(0, 1)
	for part := range result {
		exp.Add(exp, part)
	}
	return exp
}

func computeItersCount(decimalDigitsCount uint64) uint64 {
	if decimalDigitsCount == 0 {
		return 2
	}
	res := big.NewInt(10)
	for i := uint64(0); i != decimalDigitsCount; i++ {
		res.Mul(res, big.NewInt(10))
	}
	fact := big.NewInt(1)
	itersCount := uint64(1)
	for fact.Cmp(res) == -1 {
		fact.Mul(fact, big.NewInt(int64(itersCount)))
		itersCount++
	}
	return itersCount
}

func Process(decimalDigitsCount uint64, nodesCount uint64) {
	itersCount := computeItersCount(decimalDigitsCount)
	prec := uint(3.3 * float64(decimalDigitsCount))
	fmt.Printf("Iterations needed: %v\n", itersCount)
	res := delegate(itersCount, nodesCount)
	fmt.Printf("Numerator: %v\n", res.Num())
	fmt.Printf("Denumerator: %v\n", res.Denom())
	numAsFloat := new(big.Float).SetInt(res.Num())
	numAsFloat.SetPrec(prec) // Must be log_2^{10}
	denomAsFloat := new(big.Float).SetInt(res.Denom())
	denomAsFloat.SetPrec(prec)
	resAsFloat := new(big.Float).Quo(numAsFloat, denomAsFloat)
	fmt.Printf(
		"Value with %d digits after the decimal point: %s\n",
		decimalDigitsCount,
		resAsFloat.Text('f', int(decimalDigitsCount)))
}

// First 1000 digits
// 2.71828182845904523536028747135266249775724709369995957496696762772407663035
// 3547594571382178525166427427466391932003059921817413596629043572900334295260
// 5956307381323286279434907632338298807531952510190115738341879307021540891499
// 3488416750924476146066808226480016847741185374234544243710753907774499206955
// 1702761838606261331384583000752044933826560297606737113200709328709127443747
// 0472306969772093101416928368190255151086574637721112523897844250569536967707
// 8544996996794686445490598793163688923009879312773617821542499922957635148220
// 8269895193668033182528869398496465105820939239829488793320362509443117301238
// 1970684161403970198376793206832823764648042953118023287825098194558153017567
// 1736133206981125099618188159304169035159888851934580727386673858942287922849
// 9892086805825749279610484198444363463244968487560233624827041978623209002160
// 9902353043699418491463140934317381436405462531520961836908887070167683964243
// 7814059271456354906130310720851038375051011574770417189861068739696552126715
// 46889570350354
