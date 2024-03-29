package main

import (
	"fmt"
)

func main()  {
 list := []int{2,1,3,9,4,5,7,2}

 quickSort(list, 0,  len(list) - 1)

 fmt.Printf("list: %v\n", list)
}

func quickSort(arr []int, low int, high int) {
 if low < high {
  pi := partion(arr, low, high)

  // Recursively sort elements before partition and after partition
  quickSort(arr, low, pi-1)
  quickSort(arr, pi+1, high)
 }
}

func partion(arr []int, low int, high int) int {
 pivot := arr[high]
 i := low - 1

 for j := low; j < high; j++ {
  if arr[j] < pivot {
   i++

   arr[i], arr[j] = arr[j], arr[i]
  }
 }

 arr[i+1], arr[high] = arr[high], arr[i+1]
 return i + 1
}