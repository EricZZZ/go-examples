package main

import "fmt"

type Counter struct {
	Count int
}

// value receiver, does not mutate the original struct
func (c Counter) IncWrong() { c.Count++ }

// pointer receiver, mutates the original struct
func (c *Counter) IncRight() { c.Count++ }

func main() {
	c := Counter{}

	c.IncWrong()
	c.IncWrong()
	c.IncWrong()
	fmt.Println("wrong:", c.Count) // 0

	c2 := Counter{}
	c2.IncRight()
	c2.IncRight()
	c2.IncRight()
	fmt.Println("right:", c2.Count) // 3
}
