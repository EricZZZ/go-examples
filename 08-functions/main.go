package main

import (
	"errors"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	return a / b, nil
}

func apply(slice []int, fn func(int) int) {
	for i, v := range slice {
		slice[i] = fn(v)
	}
}

func double(n int) int {
	return n * 2
}

func square(n int) int {
	return n * n
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func couter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func add1(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	fmt.Println(add(2, 2))

	res, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	nums := []int{1, 2, 3, 4, 5}
	apply(nums, square)
	fmt.Println(nums)

	sum1 := sum(1, 2, 3, 4, 5)
	fmt.Println(sum1)

	greet := func() { fmt.Println("Hello") }
	greet()

	increment := couter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(add1(2, 5))
}
