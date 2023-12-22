package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	intList := []int{2,4,6,8,10}

	result := getSquereNumbersSum(intList)

	fmt.Printf("result: %v\n", result)
}

func getSquereNumbersSum(list []int) int {
	var mux sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(list))

	var sum int

	for _, v := range list {
		go increaseSum(v, &sum, &mux, &wg)
	}

	wg.Wait()

	return sum
}

func increaseSum (value int, sum *int, mux *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	squaredValue := math.Pow(float64(value), 2)

	mux.Lock()

	*sum += int(squaredValue)

	mux.Unlock()
}
