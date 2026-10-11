package main

import "fmt"

func maxFrequencyElements(nums []int) int {
	temp := make([]int, 100)
	m := 0
	ans := 0
	for _, v := range nums {
		temp[v-1]++
		m = max(m, temp[v-1])
	}

	for _, v := range temp {
		if v == m {
			ans += m
		}
	}
	return ans
}

func main() {
	fmt.Println(maxFrequencyElements([]int{1, 2, 3, 4, 5}))
}
