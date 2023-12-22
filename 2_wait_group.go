package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	intList := []int{2,4,6,8,10}

	handleSlice(intList)
}

func handleSlice(list []int) {
	var wg sync.WaitGroup

	wg.Add(len(list))

	for _, v := range list {
		go func(localValue int) {
			result := math.Pow(float64(localValue), 2)
			fmt.Printf("%v\n", result)

			wg.Done()
		} (v)
	}

	wg.Wait()
}
