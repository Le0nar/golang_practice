package main

import "fmt"

func main () {
	tempList := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempGroup := groupTempList(tempList)

	fmt.Printf("tempGroup: %v\n", tempGroup)
}

func groupTempList(list []float32) map[int][]float32 {
	// TODO: mb move step to arguments
	const step = 10
	group := make(map[int][]float32)

	for i, v := range list {
		key := int(list[i] / step) * step

		groupSlice, isSliceExist := group[key]

		if !isSliceExist  {
			group[key] = make([]float32, 1)
		}

		groupSlice = append(groupSlice, v)

		group[key] = groupSlice
	}

	return group
}
