package main

import (
 "fmt"
 "sync"
 "time"
)


func main()  {
 cnt := Counter{}

 for i := 0; i < 1000; i++ {
  go cnt.Increment()
 }

 time.Sleep(time.Second)

 value := cnt.GetValue()

 fmt.Printf("value: %v\n", value)
}

type Counter struct {
 number int
 mux sync.Mutex
}

func (c *Counter) Increment()  {
 c.mux.Lock()
 defer c.mux.Unlock()
 c.number += 1
}

func (c *Counter) GetValue() int {
 return c.number
}