package main

import (
	"fmt"
	"math/rand"
)

func main() {
	goIsAwesome := true

	if goIsAwesome {
		fmt.Println("I love go")
	}

	foods := []string{"pizza", "sushi", "tacos"}
	if choice := foods[rand.Intn(len(foods))]; choice == "pizza" {
		fmt.Println("I love pizza")
	} else {
		fmt.Println("I don't love", choice)
	}

	day := 68
	switch day {
	case 1:
		fmt.Println("It's Sunday")
	case 2:
		fmt.Println("It's Monday")
	case 3:
		fmt.Println("It's Tuesday")
	case 4:
		fmt.Println("It's Wednesday")
	case 5:
		fmt.Println("It's Thursday")
	case 6:
		fmt.Println("It's Friday")
	case 7:
		fmt.Println("It's Saturday")
	default:
		fmt.Println("It's a Invalid day")
	}

	temperature := 15
	switch {
	case temperature < 5:
		fmt.Println("It's cold")
	case temperature >= 5 && temperature < 25:
		fmt.Println("It's warm")
	case temperature >= 25:
		fmt.Println("It's hot")
	}
}
