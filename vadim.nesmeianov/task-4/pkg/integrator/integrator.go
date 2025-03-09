package integrator

import (
	"math"
	"sync"

	"github.com/golang-collections/collections/stack"
)

var SPK int
var accuracy float64
var maxTask int
var nThreads int

var wGroup sync.WaitGroup

var globalStackTaskPresent sync.Mutex

var globalStackMutex sync.Mutex
var nActive int
var globalStackPtr *stack.Stack

var sumMutex sync.Mutex
var globalSum float64

// begin, end - borders of integrating interval;
// accuracyPow - the power of accuracy (10^(accuracyPow));
// function - a mathematical function with float64 argument and float64 return value;
// n - threads quantity
func Integrate(begin, end float64, accuracyPow int, function func(float64) float64, n int) float64 {
	globalStackPtr = stack.New()
	SPK = 8
	maxTask = 10000000
	globalSum = 0
	nActive = 0
	nThreads = n
	accuracy = math.Pow(10, float64(accuracyPow))

	initTask := GetIntegrateTask(begin, end)
	globalStackPtr.Push(initTask)

	wGroup.Add(nThreads)
	for range nThreads {
		go executer()
	}
	wGroup.Wait()

	return globalSum
}

func function(x float64) float64 {
	if x == 0 {
		return 0
	} else {
		return math.Sin(1.0 / x)
	}
}

func breakCondition(sacb, sab, epsilon float64) bool {
	return !(math.Abs(sacb-sab) < epsilon)
}

func executer() {
	defer wGroup.Done()
	localStack := stack.New()
	localSum := float64(0)

	for {
		globalStackTaskPresent.Lock()
		globalStackMutex.Lock()
		currTask := globalStackPtr.Pop().(IntegrateTask)

		if globalStackPtr.Len() != 0 {
			globalStackTaskPresent.Unlock()
		}

		if !currTask.IsTerminal() {
			nActive += 1
		}

		globalStackMutex.Unlock()

		if currTask.IsTerminal() {
			break
		}

		for {
			middle := currTask.GetMiddle()
			fc := function(middle)
			fa := function(currTask.Begin)
			fb := function(currTask.End)
			sac := (fa + fc) * (middle - currTask.Begin) / 2
			scb := (fc + fb) * (currTask.End - middle) / 2
			sacb := sac + scb
			sab := (fa + fb) * (currTask.End - currTask.Begin) / 2

			if !breakCondition(sacb, sab, accuracy) {
				localSum += sacb

				if localStack.Len() == 0 {
					break
				}
				currTask = localStack.Pop().(IntegrateTask)
			} else {
				node1, node2 := currTask.SplitTask()
				localStack.Push(node1)
				currTask = node2
			}

			if (localStack.Len() > SPK) && (globalStackPtr.Len() == 0) {
				globalStackMutex.Lock()
				if globalStackPtr.Len() == 0 {
					globalStackTaskPresent.Unlock()
				}

				for (localStack.Len() != 0) && (globalStackPtr.Len() < maxTask) {
					tempTask := localStack.Pop()
					globalStackPtr.Push(tempTask)
				}

				globalStackMutex.Unlock()
			}
		}

		globalStackMutex.Lock()
		nActive -= 1

		if (nActive == 0) && (globalStackPtr.Len() == 0) {
			for range nThreads {
				terminalTask := GetTerminalTask()
				globalStackPtr.Push(terminalTask)
			}

			globalStackTaskPresent.Unlock()
		}

		globalStackMutex.Unlock()
	}

	sumMutex.Lock()
	globalSum += localSum
	sumMutex.Unlock()

}
