package main

import (
 "fmt"
 "reflect"
)

func main () {
	var val1, val2, val3, val4 = "test", 2, true, make(chan int)

	type1 := getType(val1)
	type2 := getType(val2)
	type3 := getType(val3)
	type4 := getType(val4)

	fmt.Printf("type1: %v\n", type1)
	fmt.Printf("type2: %v\n", type2)
	fmt.Printf("type3: %v\n", type3)
	fmt.Printf("type4: %v\n", type4)
}

func getType[T any](value T) string {
	parameterType := reflect.TypeOf(value)

	return parameterType.String()
}