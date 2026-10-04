package main

import "fmt"

func takeAttendance(records []int) int {
	for i, v := range records {
		if i != v {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(takeAttendance([]int{0, 1, 2, 3, 5}))
}
