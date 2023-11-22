package main

import (
	"fmt"
	"math"
)

type reactangle struct {
	a float64
	b float64
}

func (r *reactangle) getPerimeter() float64 {
	return (r.a + r.b) * 2
}

type circe struct {
	r float64
}

func (c *circe) getPerimeter() float64 {
	return 2 * math.Pi * c.r
}

type triangle struct {
	a float64
	b float64
	c float64
}

func (t *triangle) getPerimeter() float64 {
	return t.a + t.b + t.c
}

type shape interface {
	getPerimeter() float64
}

func calculateAverage(shapeList []shape) float64 {
	var result float64

	for _, v := range shapeList {
		result += v.getPerimeter()
	}

	result /= float64(len(shapeList))

	return result
}

func main () {
	list := []shape{
		&triangle{a: 1,b: 2, c: 3},
		&triangle{a: 1,b: 1, c: 1},
		&reactangle{a: 1.5,b: 1.5},
		&circe{r: 0.95541401273},
	}

	average := calculateAverage(list)

	fmt.Printf("average: %v\n", average)
}