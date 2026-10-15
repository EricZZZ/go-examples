package main

import (
	"fmt"
	"slices"
)

func lastStoneWeight(stones []int) int {
	slices.Sort(stones)
	slices.Reverse(stones)
	for len(stones) > 1 {
		k := abs(stones[0] - stones[1])
		stones = slices.Delete(stones, 0, 2)
		if k > 0 {
			stones = append(stones, k)
		}
		slices.Sort(stones)
		slices.Reverse(stones)
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(lastStoneWeight([]int{9, 3, 2, 10}))
}
