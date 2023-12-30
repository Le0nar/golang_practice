package main

import "fmt"

func main() {
	femaleNames := []string{"Саша", "Катя", "Женя", "Вика"}
	maleNames := []string{"Саша", "Витя", "Женя", "Кузя"}

	intersections := getIntersection(femaleNames, maleNames)

	fmt.Printf("intersections: %v\n", intersections)
}

// TODO: find a more optimized solution
func getIntersection[T comparable](firstList, secondList []T) []T {
	intersections := make([]T, 0)

	for _, firstListValue := range firstList {
		for _, secondListValue := range secondList {
			if firstListValue == secondListValue {
				intersections = append(intersections, secondListValue)
			}
		}
	}

	return intersections
}
