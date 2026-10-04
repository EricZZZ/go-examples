package main

import "fmt"

func maximumCount(nums []int) int {
	pos, neg := 0, 0
	for _, v := range nums {
		if v < 0 {
			neg++
		}
		if v > 0 {
			pos++
		}
	}
	return max(pos, neg)
}

func main() {
	fmt.Println(maximumCount([]int{-2, -1, -1, 1, 2, 3}))
}
