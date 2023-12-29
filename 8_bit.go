package main

import (
	"fmt"
)

func main() {
	var number int64
	number = 2

	changedValue := changeInt64Bit(number, 1, 0)

	fmt.Printf("changedValue: %v\n", changedValue)
}

func changeInt64Bit(number int64, bitValue, bitPosition int) int64 {
	if bitValue != 1 && bitValue != 0 {
		panic("Invalid value of bit")
	}

	if bitValue == 1 {
		number = number | (1 << bitPosition)
		return number
	}

	var mask int64 = ^(1 << bitPosition)

	number = number & mask
	return number
}
