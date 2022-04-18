package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := int32(6)
	staircase(n)
}

func staircase(n int32) {
	str := ""
	for i := int32(0); i < n; i++ {
		str += "#"
		fmt.Printf("%"+strconv.FormatInt(int64(n), 10)+"v\n", str)
	}
}
