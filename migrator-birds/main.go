package main

import (
	"fmt"
)

func main() {
	fmt.Println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}))
}

/*
	count repetead values
	select the lower one
*/

func migratoryBirds(arr []int32) int32 {
	bTypesCount := make([]int32, 5)

	for i := 0; i < len(arr); i++ {
		bTypesCount[arr[i]-1]++
	}

	var max = bTypesCount[0]
	var lowestType = 1

	for i := 1; i < 5; i++ {
		if bTypesCount[i] > max {
			max = bTypesCount[i]
			lowestType = i + 1
		}
	}

	return int32(lowestType)
}
