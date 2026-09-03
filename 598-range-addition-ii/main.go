package main

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		if op[0] < m {
			m = op[0]
		}
		if op[1] < n {
			n = op[1]
		}
	}
	return m * n
}

func main() {
	fmt.Println(maxCount(3, 3, [][]int{}))
}
