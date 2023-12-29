package main

import "fmt"

func main() {
	animals := []string{"cat", "cat", "dog", "cat", "tree"}

	animalSet := getItemSet(animals)

	fmt.Printf("animalSet: %v\n", animalSet)
}

func getItemSet [T comparable] (list []T) map[T]struct{} {
	set := make(map[T]struct{})

	for _, animal := range list {
		set[animal] = struct{}{}
	}

	return set
}
