package main

import "fmt"

func countOdds(low int, high int) int {
	if (high-low+1)%2 == 0 {
		return (high - low + 1) / 2
	} else {
		if low%2 == 0 {
			return (high - low + 1) / 2
		} else {
			return (high-low+1)/2 + 1
		}

	}
}

func main() {
	fmt.Println(countOdds(3, 7))
}
