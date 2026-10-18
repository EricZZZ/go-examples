package main

import "fmt"

func sumCounts(nums []int) int {
	sum := 0
	for i := range len(nums) {
		set := make(map[int]struct{})
		for j := i; j < len(nums); j++ {
			set[nums[j]] = struct{}{}
			x := len(set)
			sum += (x * x)
		}
	}
	return sum
}

func main() {
	fmt.Println(sumCounts([]int{1, 2, 1}))
}
