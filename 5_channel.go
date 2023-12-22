package main

import (
	"fmt"
	"time"
)

const lifetime = 5

func main() {
	ch := make(chan int)

	go startWriter(ch)
	go startReader(ch)

	time.Sleep(time.Second * lifetime)
}

func startWriter(ch chan int) {
	var value int

	for  {
		time.Sleep(time.Millisecond * 500)

		fmt.Printf("Write: %d \n", value)
		ch <- value


		value ++
	}
}

func startReader(ch chan int) {
	for  {
		value := <- ch

		fmt.Printf("Read: %d \n", value)
	}
}
