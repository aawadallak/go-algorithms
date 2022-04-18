package main

import (
	"fmt"
)

func main() {
	arr := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	fmt.Println(diagonalDifference(arr))
}

func diagonalDifference(arr [][]int32) int32 {
	var left, right int32

	for i := 0; i < len(arr); i++ {
		left += arr[i][i]
		right += arr[i][len(arr)-i-1]
	}

	if left < right {
		return right - left
	}

	return left - right
}
