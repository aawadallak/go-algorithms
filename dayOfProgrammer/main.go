package main

import (
	"fmt"
)

func main() {
	// fmt.Println(dayOfProgrammer(1916))
	// fmt.Println(dayOfProgrammer(2016))
	// fmt.Println(dayOfProgrammer(2017))
	fmt.Println(dayOfProgrammer(1918))
}

func dayOfProgrammer(year int32) string {
	var days int32 = 243
	if year > 1918 {
		if year%400 == 0 || (year%4 == 00 && year%100 != 0) {
			days += 1
		}
	} else if year == 1918 {
		days -= 13
	} else {
		if year%4 == 0 {
			days += 1
		}
	}

	date := 256 - days
	return fmt.Sprintf("%02d.%02d.%04d", date, 9, year)
}
