package main

import (
	"fmt"
)

func main() {
	fmt.Println(kangaroo(2564, 5393, 5121, 2836))
	fmt.Println(kangaroo(14, 4, 98, 2))
	fmt.Println(kangaroo(21, 6, 47, 3)) // execpted no
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 > x2 || v1 > v2 {
		for x1 < x2 || x2 <= 10000 {
			if (x1 + v1) == (x2 + v2) {
				return "YES"
			}
			x1 += v1
			x2 += v2
		}
	}

	return "NO"
}
