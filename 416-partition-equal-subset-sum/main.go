package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	t := sum / 2
	dp := make([]bool, t+1)
	dp[0] = true
	for _, v := range nums {
		for i := t; i >= v; i-- {
			if dp[i-v] {
				dp[i] = true
			}
		}
	}
	return dp[t]
}

func main() {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
