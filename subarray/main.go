package main

import (
	"fmt"
)

func main() {
	fmt.Println(birthday([]int32{2, 2, 1, 3, 2}, 4, 2))
}

/*

 */

func birthday(s []int32, d int32, m int32) int32 {
	var count int32

	for i := int32(0); i < int32(len(s)); i++ {
		if i+m > int32(len(s)) {
			continue
		}

		var match int32
		for j := 0; j < len(s[i:i+m]); j++ {
			match += s[i : i+m][j]
		}

		if match == d {
			count++
		}
	}
	return count
}
