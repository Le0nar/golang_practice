package main

import "fmt"

func calculate(ch chan int) {
	var value int

	for i := 0; i < 100; i++ {
		value++
	}

	ch <- value
}

func main() {
	ch := make(chan int)
	go calculate(ch)
	calculatedValue := <-ch

	fmt.Printf("calculatedValue: %v\n", calculatedValue)
}
