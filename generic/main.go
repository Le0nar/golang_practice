package main

import "fmt"

func main() {
	intList := []int{1, 2, 3, 4, 5}
	floatList := []float64{1, 2, 3, 4, 5}

	changedList := handleList(intList)

	fmt.Printf("changedList: %v\n", changedList)

	changedList2 := handleList(floatList)

	fmt.Printf("changedList: %v\n", changedList2)
}

type Number interface {
	int | float64
}

func handleList[T Number](list []T) []T {
	for i, v := range list {
		list[i] = v  + 10
	}
	return list
}