package main

import (
	"fmt"
	"math"
)

func main() {
	grades := []int32{4, 73, 67, 38, 33}
	fmt.Println(gradingStudents(grades))
}

func gradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		grt := int32(math.Ceil(float64(grades[i])/5) * 5)
		if grades[i] >= 38 && (grt-grades[i]) < 3 {
			grades[i] = grt
		}
	}

	return grades
}
