package main

import "fmt"

func main() {
	fmt.Print("Hello!\n")
	fmt.Println("Go is awesome")

	name := "eric"
	count := 2
	fmt.Printf("My name is %s and count is %d\n", name, count)

	sentence := fmt.Sprintf("My name is %s and count is %d", name, count)
	fmt.Printf("Sentence is: %s\n", sentence)

	var age int

	fmt.Print("Enter your name:")
	fmt.Scan(&name)

	fmt.Print("Enter your age:")
	fmt.Scan(&age)

	fmt.Printf("Hello,%s! you are %d years old.\n", name, age)

	var favColor1 string
	var favColor2 string

	fmt.Print("Enter your two favorite colors: ")
	fmt.Scan(&favColor1, &favColor2)

	fmt.Println("Favorite colors ", favColor1, favColor2)
}
