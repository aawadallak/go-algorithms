package main

import "fmt"

var arr = []int32{-4, 3, -9, 0, 4, 1}

func main() {
	plusMinus(arr)
}

func plusMinus(arr []int32) {
	n := float64(len(arr))
	var pos, neg, zero float64

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zero += 1
			continue
		}

		if arr[i] < 0 {
			neg += 1
		} else {
			pos += 1
		}
	}

	fmt.Printf("%.6f\n", pos/n)
	fmt.Printf("%.6f\n", neg/n)
	fmt.Printf("%.6f\n", zero/n)
}
