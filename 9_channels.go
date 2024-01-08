package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	originList := []int{1, 2,3,4,5,6,7,8,9,10}
	firstCh := make(chan int)
	secondCh := make(chan int)

	go start(firstCh, secondCh, len(originList), &wg)

	go func ()  {
		for _, v := range originList {
			firstCh <- v
		}
		wg.Done()
	}()

	go func ()  {
		for i := 0; i < len(originList); i++ {
			value := <- secondCh
			fmt.Printf("value: %v\n", value)
		}
		wg.Done()
	}()

	wg.Wait()
}

// TODO: add closing for goroutine
// TODO: mb make first chan only for read and second only for write
func start(readCh, writeCh chan int, listLength int, wg *sync.WaitGroup) {
	for i := 0; i < listLength; i++ {
		value := <- readCh
		value *= 2
		writeCh <- value
	}

	wg.Done()
}
