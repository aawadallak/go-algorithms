package main

import (
	"fmt"
)

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
}

func breakingRecords(scores []int32) []int32 {
	var max, min = scores[0], scores[0]
	records := []int32{0, 0}

	for i := 0; i < len(scores); i++ {
		if scores[i] < min {
			min = scores[i]
			records[1]++
		}

		if scores[i] > max {
			max = scores[i]
			records[0]++
		}
	}

	return records
}
