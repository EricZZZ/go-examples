package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------------------")

	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------------------")

	for i := range 5 {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println("-----------------------")

	languages := []string{"Go", "Python", "Java", "C++", "JavaScript"}
	for i, lang := range languages {
		fmt.Println(i, lang)
	}
}
