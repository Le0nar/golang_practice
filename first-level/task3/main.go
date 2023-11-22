package main

import "fmt"

func doSmth(list []int) {
	result := make(map[int]int)
	for _, v := range list {
		result[v % 5] = v
	}

	fmt.Printf("result: %v\n", result)
}

func main() {
	doSmth([]int{8, 9})
}
