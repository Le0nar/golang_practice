package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = b, a

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}
