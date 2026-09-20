package main

import "fmt"

func reverseDegree(s string) int {
	res := 0

	for i, c := range s {
		res += (26 - (int(c) - 97)) * (i + 1)
	}

	return res
}

func main() {
	fmt.Println(reverseDegree("abc"))
}
