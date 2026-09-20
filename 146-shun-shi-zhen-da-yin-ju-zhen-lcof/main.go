package main

import "fmt"

func spiralArray(array [][]int) []int {
	if len(array) == 0 || len(array[0]) == 0 {
		return []int{}
	}

	res := make([]int, 0)

	top := 0
	bottom := len(array) - 1
	left := 0
	right := len(array[0]) - 1

	for top <= bottom && left <= right {

		for i := left; i <= right; i++ {
			res = append(res, array[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, array[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, array[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, array[i][left])
			}
			left++
		}
	}

	return res
}

func main() {
	fmt.Println(spiralArray([][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}))
}
