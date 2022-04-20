package main

import (
	"fmt"
)

// count only fruits in range between s and t

func countApplesAndOranges(
	s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var apple, orange int32

	for i := 0; i < len(apples); i++ {
		x := a + apples[i]
		if x >= s && x <= t {
			apple++
		}
	}

	for i := 0; i < len(oranges); i++ {
		y := b + oranges[i]
		if y >= s && y <= t {
			orange++
		}
	}

	fmt.Printf("%d\n", apple)
	fmt.Printf("%d\n", orange)
}

func main() {
	countApplesAndOranges(
		7, 11, 5, 15,
		[]int32{-2, 2, 1},
		[]int32{5, -6},
	)
}
