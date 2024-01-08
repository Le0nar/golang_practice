package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const workersQuantity = 5

func main() {
	quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	intChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(workersQuantity)

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < workersQuantity; i++ {
		go startWorker(ctx, &wg, intChan, i)
	}

	go seedRandInt(intChan)

	<-quit
	cancel()

	wg.Wait()

	fmt.Printf("The program exits")
}

func startWorker(ctx context.Context, wg *sync.WaitGroup, ch chan int, workerId int) {
	for {
		select {
		case value := <- ch:
			fmt.Printf("worker %d say: %v\n",workerId, value)
			
		case <- ctx.Done():
			fmt.Printf("Worker %d finishes work \n", workerId)
			wg.Done()
			return
		}
	}
}

func seedRandInt(ch chan int) {
	for {
		// TODO: timeout is optional
		time.Sleep(time.Millisecond * 500)

		ch <- rand.Intn(100)
	}
}
