package office

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type AcRequest struct {
	temp int
	op   Operand
}

type ac struct {
	maxTemp     int
	minTemp     int
	currentTemp int
	requests    []AcRequest
}

type Ac interface {
	NewRequest(temp int, op Operand) error
	GetTemperature() int
}

func (instance *ac) processRequest(req AcRequest) error {
	instance.requests = append(instance.requests, req)

	if instance.currentTemp != -1 {
		minNewTemp := instance.minTemp
		maxNewTemp := instance.maxTemp

		for _, rq := range instance.requests {
			switch rq.op {
			case BiggerOrEqual:
				if rq.temp >= minNewTemp {
					minNewTemp = rq.temp
				}

			case LessOrEqual:
				if rq.temp <= maxNewTemp {
					maxNewTemp = rq.temp
				}

			default:
				return errors.New("unknown operator")
			}
		}

		if minNewTemp <= maxNewTemp {
			instance.currentTemp = minNewTemp
		} else {
			instance.currentTemp = -1
		}
	}

	return nil
}

// TODO: implement
func (instance *ac) NewRequest(temp int, op Operand) error {
	req := AcRequest{temp, op}
	return instance.processRequest(req)
}

func (instance ac) GetTemperature() int {
	return instance.currentTemp
}

func GetAc() Ac {
	const minTemp = 15
	const maxTemp = 30

	instance := new(ac)
	instance.maxTemp = maxTemp
	instance.minTemp = minTemp
	instance.currentTemp = 0
	instance.requests = make([]AcRequest, 0, 2)
	instance.requests = append(instance.requests, AcRequest{minTemp, BiggerOrEqual})
	instance.requests = append(instance.requests, AcRequest{maxTemp, LessOrEqual})

	return instance
}

type department struct {
	depId         int
	employeesNum  int
	employeesCame int
	ac            Ac
}

type Department interface {
	EventEmployeeCame(temp int, op Operand) error
	SetCapacity(int)
	GetTemperature() int
}

func (instance department) GetTemperature() int {
	return instance.ac.GetTemperature()
}

func (instance *department) EventEmployeeCame(temp int, op Operand) error {
	if instance.employeesCame >= instance.employeesNum {
		errStr := fmt.Sprintf("department %d is overcrowded", instance.depId)
		return errors.New(errStr)
	}

	return instance.ac.NewRequest(temp, op)
}

func (instance *department) SetCapacity(cap int) {
	instance.employeesNum = cap
}

func GetDepartment(id int) Department {
	instance := new(department)
	instance.depId = id
	instance.employeesCame = 0
	instance.ac = GetAc()

	return instance
}

type office struct {
	depNum int
	dep    []Department
}

type Office interface {
	EventEmployeeCame(int, int, Operand) error
	SetDepartmentCapasity(int, int) error

	GetTempInDepartment(int) int
}

type Operand int

const (
	BiggerOrEqual Operand = iota
	LessOrEqual
	OperandFailed
)

func OperandFromString(str string) (Operand, error) {
	if str == ">=" {
		return BiggerOrEqual, nil
	}
	if str == "<=" {
		return LessOrEqual, nil
	}

	errStr := fmt.Sprintf("failed to parse operand %s", str)
	return OperandFailed, errors.New(errStr)
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Compare[T Ordered](op Operand, a1, a2 T) (bool, error) {
	switch op {
	case BiggerOrEqual:
		return a1 >= a2, nil
	case LessOrEqual:
		return a1 <= a2, nil
	default:
		return false, errors.New("unknown operand")
	}
}

func (instance office) GetTempInDepartment(id int) int {
	return instance.dep[id].GetTemperature()
}

func (instance *office) EventEmployeeCame(dep, temp int, op Operand) error {
	if dep < instance.depNum {
		return instance.dep[dep].EventEmployeeCame(temp, op)
	}

	errStr := fmt.Sprintf("dep_id = %d is too big", dep)
	return errors.New(errStr)
}

func (instance *office) SetDepartmentCapasity(num, size int) error {
	if num < instance.depNum {
		instance.dep[num].SetCapacity(size)
		return nil
	}

	errStr := fmt.Sprintf("dep_id = %d is too big", num)
	return errors.New(errStr)
}

func GetOffice(n int) Office {
	instance := new(office)
	instance.depNum = n
	instance.dep = make([]Department, n)

	for i, _ := range instance.dep {
		instance.dep[i] = GetDepartment(i)
	}

	return instance
}

func Run(logOutput bool) {
	var n int
	err := readValue(&n, "Enter N:", logOutput)
	if err != nil {
		handleError(err)
	}

	_office := GetOffice(n)

	for i := 0; i < n; i += 1 {
		var k int
		msgStr := fmt.Sprintf("Enter K for department %d", i)
		err = readValue(&k, msgStr, logOutput)
		if err != nil {
			handleError(err)
		}

		_office.SetDepartmentCapasity(i, k)

		for j := 0; j < k; j += 1 {
			op, temp, err := readRequest()
			if err != nil {
				handleError(err)
			}

			_office.EventEmployeeCame(i, temp, op)

			currTemp := _office.GetTempInDepartment(i)
			fmt.Printf("%d\n", currTemp)
		}
	}
}

func readRequest() (Operand, int, error) {
	var str string
	var k int
	_, err := fmt.Scanf("%s%d\n", &str, &k)
	if err != nil {
		return OperandFailed, 0, err
	}

	lexemsArr := strings.Fields(str)
	if len(lexemsArr) != 1 {
		errStr := fmt.Sprintf("the string: \" %s \" has %d lexems, must be 2", str, len(lexemsArr))
		return OperandFailed, 0, errors.New(errStr)
	}
	op, err := OperandFromString(lexemsArr[0])
	if err != nil {
		return OperandFailed, 0, err
	}

	return op, k, nil
}

func readValue[T any](x *T, msg string, logOutput bool) error {
	if logOutput {
		fmt.Println(msg)
	}

	_, err := fmt.Scan(x)
	return err
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
