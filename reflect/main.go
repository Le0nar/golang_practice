package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Ticket struct {
		Number int
		City   string
	}

	ticket := Ticket{
		Number: 333333,
		City:   "Izhevsk",
	}

	value := getValue(ticket, "City")

	fmt.Printf("value: %v\n", value)

}

func getValue[T any](structure T, fieldName string) any {
	refValue := reflect.ValueOf(structure)

	value := refValue.FieldByName(fieldName)

	return value
}