package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int
	path := []int{}

	var backtrack func(start int, target int)

	backtrack = func(start, target int) {
		if target == 0 {
			result = append(result, append([]int{}, path...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}

			path = append(path, candidates[i])

			backtrack(i, target-candidates[i])

			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)

	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
