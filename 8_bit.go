package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int64
	number = 20
	value, index := 0, 2

	changedValue := changeInt64(number, value, index)

	fmt.Printf("changedValue: %d\n", changedValue)
}

func changeInt64(number int64, value, index int) int64 {
	binaryValue := strconv.FormatInt(number, 2)

	stringedValue := strconv.Itoa(value)

	// 1) Cut first part of the string (from 0 to index)
	// 2) Add new value
	// 3) Cut rest of the string (from index + 1 to the end)
	binaryValue = binaryValue[:index] + stringedValue + binaryValue[index + 1:]

	changedValue, err := strconv.ParseInt(binaryValue, 2, 64)
	if err != nil {
		panic(err)
	}

	return changedValue
}
