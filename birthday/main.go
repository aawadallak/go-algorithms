package main

import (
	"fmt"
)

func main() {
	candles := []int32{44, 53, 31, 27, 77, 60, 66, 77, 26, 36}
	fmt.Println(birthdayCakeCandles(candles))
}

func birthdayCakeCandles(candles []int32) int32 {
	var max, count int32
	var values []int32

	for i := 0; i < len(candles); i++ {
		values = append(values, candles[i])
		if candles[i] > max {
			max = candles[i]
		}
	}

	for i := 0; i < len(values); i++ {
		if values[i] == max {
			count++
		}
	}

	return count
}
