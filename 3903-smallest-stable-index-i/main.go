package main

import (
	"fmt"
)

func firstStableIndex(nums []int, k int) int {
	// for i, _ := range nums {
	// 	max := slices.Max(nums[0 : i+1])
	// 	min := slices.Min(nums[i:len(nums)])
	// 	if max-min <= k {
	// 		return i
	// 	}
	// }
	// return -1
	n := len(nums)

	prefixMax := make([]int, n)
	suffixMin := make([]int, n)

	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(prefixMax[i-1], nums[i])
	}

	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], nums[i])
	}

	for i, _ := range nums {
		if prefixMax[i]-suffixMin[i] <= k {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstStableIndex([]int{5, 0, 1, 4}, 3))
}
