package main

import "fmt"

func main() {
	originList := []int{1, 2,3,4,5,6,7,8,9,10}

	firstCh := make(chan int)

	bufSize := len(originList) - 1
	secondCh := make(chan int, bufSize)

	go start(firstCh, secondCh)

	for _, v := range originList {
		firstCh <- v
	}

	for i := 0; i < len(originList); i++ {
		select {
		case value := <- secondCh:
		fmt.Printf("value: %v\n", value)
		}
	}
}

// TODO: add closing for goroutine
// TODO: mb make first chan only for read and second only for write
func start(readCh, writeCh chan int) {
	for {
		select {
		case value := <- readCh:
		value *= 2
		writeCh <- value
		}
	}
}