package main

import (
	"fmt"
	"math"
	"sync"
	"time"

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

func function(x float64) float64 {
	return math.Cos(1.0 / x)
}

type IntegrateTask struct {
	Begin    float64
	End      float64
	terminal bool
}

func GetTerminalTask() IntegrateTask {
	return IntegrateTask{
		Begin:    0,
		End:      0,
		terminal: true,
	}
}

func GetIntegrateTask(a, b float64) IntegrateTask {
	return IntegrateTask{
		Begin:    a,
		End:      b,
		terminal: false,
	}
}

func (task IntegrateTask) IsTerminal() bool {
	return task.terminal
}

func (task IntegrateTask) GetMiddle() float64 {
	return (task.Begin + task.End) / 2
}

func (task IntegrateTask) SplitTask() (IntegrateTask, IntegrateTask) {
	middle := task.GetMiddle()

	task1 := IntegrateTask{
		Begin:    task.Begin,
		End:      middle,
		terminal: false,
	}

	task2 := IntegrateTask{
		Begin:    middle,
		End:      task.End,
		terminal: false,
	}

	return task1, task2
}

func breakCondition(sacb, sab, epsilon float64) bool {
	return !(math.Abs(sacb-sab) < epsilon)
}

func executer(id int) {
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
				if globalStackPtr.Len() == 0 { // TODO
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

func main() {
	globalStackPtr = stack.New()
	SPK = 8
	maxTask = 10000000
	accuracy = 1e-16
	nThreads = 8

	globalSum = 0
	nActive = 0

	wGroup.Add(nThreads)

	a := float64(0.005)
	b := float64(40000000)

	initTask := GetIntegrateTask(a, b)

	globalStackPtr.Push(initTask)

	start := time.Now()
	for i := range nThreads {
		go executer(i)
	}

	wGroup.Wait()
	end := time.Now()
	fmt.Println("S = ", globalSum, " from ", a, " to ", b)
	fmt.Println("Time elapsed: ", end.Sub(start))
	fmt.Println("Threads: ", nThreads)
}
