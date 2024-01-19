package main

import (
	"errors"
	"fmt"
)

func main()  {
 list := []int{-3333,-29,3,9,12, 24}

 index, err := searchIndex(list, -29)
 if err != nil {
  fmt.Printf("%s \n", err.Error())
  return
 }

 fmt.Printf("index: %v\n", index)
}

func searchIndex(list []int, value int) (int, error) {
 leftIndex := 0
 rightIndex := len(list) - 1
 
 searchedValueIndex := searchBinary(list, value, leftIndex, rightIndex)

 if searchedValueIndex == -1 {
  return -1, errors.New("Number is missing")
 }

 return searchedValueIndex, nil
}

// TODO: rename
func searchBinary (list []int, value, leftIndex, rightIndex  int) int {
 if leftIndex == rightIndex && list[leftIndex] == value {
  return leftIndex
 }

 middleIndex := (leftIndex + rightIndex) / 2
 middleIndexValue := list[middleIndex]

 if middleIndexValue == value {
  return middleIndex
 }

 if  value < middleIndexValue {
  return searchBinary(list, value, leftIndex, middleIndex)
 }

 if value > middleIndexValue {
  return searchBinary(list, value, middleIndex + 1, rightIndex)
 }

 return -1
}