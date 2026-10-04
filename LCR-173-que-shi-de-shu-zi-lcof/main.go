package main

import "fmt"

func takeAttendance(records []int) int {
	for i, v := range records {
		if i != v {
			return i
		}
	}
	return len(records)
}

func takeAttendance2(records []int) int {
	i := 0
	j := len(records) - 1
	for i <= j {
		m := (i + j) / 2
		if records[m] == m {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return i
}

func main() {
	fmt.Println(takeAttendance2([]int{0, 1, 2, 3, 5}))
}
