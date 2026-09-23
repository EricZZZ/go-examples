package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	pre := 0
	m := nums[0]
	for _, v := range nums {
		pre = max(pre+v, v)
		m = max(pre, m)
	}
	return m
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
