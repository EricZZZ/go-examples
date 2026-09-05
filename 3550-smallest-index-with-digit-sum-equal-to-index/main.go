package main

import "fmt"

func smallestIndex(nums []int) int {
	digitSum := func(n int) int {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		return sum
	}
	for i, num := range nums {
		if digitSum(num) == i {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(smallestIndex([]int{1, 3, 2}))
}
