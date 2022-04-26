package main

import (
	"fmt"
)

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var pairs int32

	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n; j++ {
			if i < j && (ar[i]+ar[j])%k == 0 {
				pairs += 1
			}
		}
	}

	return pairs
}
