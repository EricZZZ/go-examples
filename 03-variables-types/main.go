package main

import "fmt"

func main() {
	var animal1 string = "cat"
	var animal2 = "dog"
	animal3 := "monkey"
	var animal4 string
	animal4 = "lion"

	var defaultNum int
	var defaultBool bool
	var defaultStr string

	fmt.Println(animal1, animal2, animal3, animal4)
	fmt.Println(defaultNum, defaultBool, defaultStr)

	var (
		name    string = "John"
		age     int    = 33
		address string = "22414 go street"
	)

	fmt.Println(name, age, address)

	const pi = 3.14
	const port = 8080

	var num1 int = 1990
	var num2 int64 = 123456789
	var num3 int8 = 127
	var num4 uint = 100
	fmt.Println(num1, num2, num3, num4)

	var amount1 float32 = 13.99
	var amount2 float64 = 88.88
	amount3 := 829.99
	fmt.Println(amount1, amount2, amount3)

	var isTrue bool = true
	isFalse := false

	fmt.Println(isTrue, isFalse)
}
