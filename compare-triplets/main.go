package main

import (
	"fmt"
)

func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}

	out := compareTriplets(a, b)
	fmt.Println(out)
}

func compareTriplets(a []int32, b []int32) []int32 {
	scores := []int32{0, 0}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			scores[0] += 1
		} else {
			scores[1] += 1
		}
	}

	return scores
}
