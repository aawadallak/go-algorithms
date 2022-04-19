package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(timeConversion("09:05:45PM"))
}

func timeConversion(s string) string {
	n := len(s) - 2
	suffix := s[n:]
	arr := strings.SplitN(s[:n], ":", -1)

	hour, err := strconv.Atoi(arr[0])
	if err != nil {
		return ""
	}

	if hour == 12 {
		hour = 0
	}

	if suffix == "PM" {
		hour += 12
	}

	return fmt.Sprintf("%02d:%s:%s", hour, arr[1], arr[2])
}
