package main

import "fmt"

// TODO: fix it

var justString string

func someFunc () {
	v := "aasaasddasdasdasasasdaasasadsdasad"

	justString := v[:5]
	fmt.Printf("justString: %v\n", justString)
}

func main () {
	someFunc()
}