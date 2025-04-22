package transformer

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"task-3/internal/data"
)

func transformElem(raw data.RawDataElement) (data.DataElement, error) {
	var elem data.DataElement
	var NumCode int
	var err error
	if raw.NumCode == "" {
		NumCode = 0
	} else {
		NumCode, err = strconv.Atoi(raw.NumCode)
		if err != nil {
			return data.DataElement{}, fmt.Errorf("fail to convert NumCode: %v", err)
		}
	}
	elem.NumCode = NumCode
	// if len(raw.CharCode) != 3 {
	// 	err = fmt.Errorf("incorrect length of CharCode: %v", raw.CharCode)
	// 	return data.DataElement{}, err
	// }
	elem.CharCode = raw.CharCode
	ValueStr := raw.Value
	ValueStr = strings.Replace(ValueStr, ",", ".", 1)
	Value, err := strconv.ParseFloat(ValueStr, 64)
	if err != nil {
		return data.DataElement{}, fmt.Errorf("fail to convert Value: %v", ValueStr)
	}
	elem.Value = Value
	return elem, nil
}

func Transform(rawElements data.RawDataElements) (data.DataElements, error) {
	var elements data.DataElements
	for _, rawElement := range rawElements {
		element, err := transformElem(rawElement)
		if err != nil {
			return data.DataElements{}, err
		}
		elements = append(elements, element)
	}
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Value > elements[j].Value
	})
	return elements, nil
}
