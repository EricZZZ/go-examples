package main

import (
	"fmt"
)

func countKDifference(nums []int, k int) int {
	ans := 0
	for i := range len(nums) - 1 {
		j := i + 1
		for j < len(nums) {
			if abs(nums[i]-nums[j]) == k {
				ans++
			}
			j++
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(countKDifference([]int{1, 2, 2, 1}, 1))
}
