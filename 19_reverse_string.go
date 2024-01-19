package main

import "fmt"

func main()  {
 const value = "1a2b3c4d5e6f 7x8x9x0 日本語"

 reversedValue := reverseString(value)

 fmt.Printf("reversedValue: %v\n", reversedValue)
}

func reverseString (value string) string {
 var reversedValue string
 for _, symbol := range value {
  reversedValue = string(symbol) + reversedValue
 }

 return reversedValue
}