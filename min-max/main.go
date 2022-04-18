package main

import (
	"fmt"
	"sort"
)

func main() {
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
}

func miniMaxSum(arr []int32) {
	var min, max int64

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr); i++ {
		index := int64(arr[i])

		if i == 0 {
			min = index
			continue
		}

		if i == len(arr)-1 {
			max += index
			continue
		}

		min += index
		max += index
	}

	fmt.Printf("%d %d\n", min, max)
}
