package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Stopping goroutine by context
	ctx, cancel := context.WithCancel(context.Background())
	go doFirstJob(ctx)
	go cancelFirstJob(cancel)

	// Stopping goroutine by channel
	ch := make(chan struct{})
	go doSecondJob(ch)
	go cancelSecondJob(ch)

	for {}
}

func doFirstJob(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("First goroutine was finished by context \n")
			return

		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Doing first job... \n")
		}
	}
}

func cancelFirstJob(cancel context.CancelFunc) {
	time.Sleep(time.Second * 3)
	cancel()
}

func doSecondJob(ch chan struct{})  {
	for {
		select {
		case <-ch:
			fmt.Printf("Second goroutine was finished by channel \n")
			close(ch)
			return
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Doing second job... \n")
		}
	}
}

func cancelSecondJob(ch chan struct{}) {
	time.Sleep(time.Second * 5)
	ch <- struct{}{}
}
