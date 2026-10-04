package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := range numRows {
		row := []int{}

		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				val := res[i-1][j-1] + res[i-1][j]
				row = append(row, val)
			}
		}

		res = append(res, row)
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
