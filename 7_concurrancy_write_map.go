package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var mux sync.Mutex
	
	object := make(map[int]int)

	ch := make(chan [2]int)

	go seedRand(ch, 0)
	go seedRand(ch, 1)

	for v := range ch {
		key, value := v[0], v[1]

		mux.Lock()
		object[key] = value
		mux.Unlock()

		fmt.Printf("object: %v\n", object)
	}

}

func seedRand(ch chan [2]int, value int) {
	for {
		// timeout is optional
		time.Sleep(time.Millisecond * 500)

		key := rand.Intn(10)
		ch <- [2]int{key,value}
	}
}
