package main

import "fmt"

func sumOfSquares(nums []int) int {
	ans := 0
	n := len(nums)
	for i, v := range nums {
		if n%(i+1) == 0 {
			ans += (v * v)
		}
	}
	return ans
}

func main() {
	fmt.Println(sumOfSquares([]int{1, 2, 3, 4}))
}
