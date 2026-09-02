package main

import "fmt"

type Number interface {
	IsEven() bool
	IsOdd() bool
}

type MyInt int

func (i MyInt) IsEven() bool {
	return i%2 == 0
}

type Person struct {
	Name       string
	age        int
	intersests [2]string
}

func main() {
	var i MyInt = 42
	fmt.Println(i.IsEven())

	var p1 any
	var p2 any

	p1 = Person{Name: "Alice", age: 30, intersests: [2]string{"reading", "hiking"}}
	p2 = Person{Name: "Alice", age: 30, intersests: [2]string{"coding", "gaming"}}

	fmt.Println(p1 == p2)
}
