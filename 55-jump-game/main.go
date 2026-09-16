package main

import "fmt"

func canJump(nums []int) bool {
	s := 0
	for i := 0; i <= len(nums); i++ {
		if i > s {
			return false
		}

		s = max(s, i+nums[i])

		if s >= len(nums)-1 {
			return true
		}
	}

	return true
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
