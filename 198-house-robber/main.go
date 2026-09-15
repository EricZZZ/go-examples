package main

import "fmt"

func rob(nums []int) int {
	prev2 := 0
	prev1 := 0
	for _, v := range nums {
		// dp[i]=max(dp[i−1],dp[i−2]+nums[i])
		current := max(prev1, prev2+v)
		prev2 = prev1
		prev1 = current

	}
	return prev1
}

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
