package main

import "fmt"

type Crane interface {
	JackUp() string
	Hoist() string
}

type CraneA struct {
	work int
}

func (c CraneA) Work() {
	fmt.Println("using A")
}

func (c CraneA) JackUp() string {
	c.Work()
	return "jackup"
}

func (c CraneA) Hoist() string {
	c.Work()
	return "hoist"
}

type CraneB struct {
	boot string
}

func (c CraneB) Boot() {
	fmt.Println("using B")
}

func (c CraneB) JackUp() string {
	c.Boot()
	return "jackup"
}

func (c CraneB) Hoist() string {
	c.Boot()
	return "hoist"
}

type ConstructionCompany struct {
	Crane Crane
}

func (c *ConstructionCompany) Build() {
	fmt.Println(c.Crane.JackUp())
	fmt.Println(c.Crane.Hoist())
	fmt.Println("complete")
}

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

type Humam interface {
	Say(string) string
	Walk(int)
}

type Man interface {
	Exercise()
	// Say(string, string) string
	Humam
}

func (p Person) Say(s1 string) string {
	return s1
}

func (p Person) Walk(distance int) {
	fmt.Println(distance)
}

func (p Person) Exercise() {
	fmt.Println("exercise")
}

func main() {
	var i MyInt = 42
	fmt.Println(i.IsEven())

	var p1 any
	var p2 any

	p1 = Person{Name: "Alice", age: 30, intersests: [2]string{"reading", "hiking"}}
	p2 = Person{Name: "Alice", age: 30, intersests: [2]string{"coding", "gaming"}}

	fmt.Println(p1 == p2)

	fmt.Println()
	// using CraneA
	company := ConstructionCompany{CraneA{}}
	company.Build()
	fmt.Println()
	// using CraneB
	company.Crane = CraneB{}
	company.Build()

	fmt.Println()

	var anything any

	anything = 1
	println(anything)
	fmt.Println(anything)

	anything = "something"
	println(anything)
	fmt.Println(anything)

	anything = complex(1, 2)
	println(anything)
	fmt.Println(anything)

	anything = 1.2
	println(anything)
	fmt.Println(anything)

	anything = []int{}
	println(anything)
	fmt.Println(anything)

	anything = map[string]int{}
	println(anything)
	fmt.Println(anything)

	fmt.Println()

	var h Man = Person{Name: "h", age: 0, intersests: [2]string{}}
	fmt.Println(h.Say("hello"))
}
