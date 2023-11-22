package main

import "fmt"

type myFloat struct {
	value float64
}

func main() {
	test := myFloat{value: 10}

	test.subtract(9).multiply(5).multiply(999.23).multiply(0).subtract(1)

	fmt.Println(test.value)
}

func (m *myFloat) subtract(divider float64) *myFloat {
	m.value -= divider

	return m
}

func (m *myFloat) multiply(multiplier float64) *myFloat {
	m.value *= multiplier
	return m
}
