package main

import (
	"fmt"
	"slices"
)

func main() {
	var zeroArray [5]int
	var myArray = [3]int{5, 10, 15}
	var cars = [5]string{1: "Volvo", 4: "BMW"}
	var snacks = [...]string{"chips", "popcorn", "peanuts"}

	fmt.Println(zeroArray)
	fmt.Println(myArray)
	fmt.Println(cars)
	fmt.Println(snacks)

	fmt.Println(len(snacks))

	// compare arrays
	x := [3]int{1, 2, 3}
	y := [3]int{1, 2, 3}
	z := [...]int{1, 2, 2}

	fmt.Println(x == y)
	fmt.Println(x == z)

	someSlice := []int{5, 10, 15}
	var carSlice = []string{1: "volvo", 4: "BMW"}
	var sliceWithNoValue []int

	fmt.Println(someSlice)
	fmt.Println(carSlice)
	fmt.Println(sliceWithNoValue)

	// compare slices
	fruits1 := []string{"apple", "banana", "peach"}
	fruits2 := []string{"apple", "banana", "peach"}
	fmt.Println(slices.Equal(fruits1, fruits2))

	fruits1 = append(fruits1, "orange")
	fmt.Println(slices.Equal(fruits1, fruits2))

	fruits1 = append(fruits1, fruits2...)
	fmt.Println(fruits1)

	nums2 := make([]int, 0, 20)
	nums2 = append(nums2, 2, 10, 55, 23)
	fmt.Println(nums2, len(nums2), cap(nums2))

	// slice is sharing, not copying
	nums3 := nums2[:4]
	nums2[0] = 10
	fmt.Println(nums2[0], nums3[0])

	// delete slice
	s := []int{1, 2, 3, 4}
	s = append(s[:0], s[0+1:]...)
	fmt.Println(s)

	// copy slice
	values := []int{5, 6, 7, 8}
	newValues := make([]int, 4)
	copy(newValues, values)
	newValues[0] = 1222
	fmt.Println(values)
	fmt.Println(newValues)
}
