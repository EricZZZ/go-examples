package main

import "fmt"

func main() {
	prices := map[string]int{"apple": 12, "banana": 22}

	// No
	cost := prices["mango"]
	fmt.Println(cost) // 0

	// Yes
	cost, ok := prices["mango"]
	if !ok {
		fmt.Println("unknown product: mango")
		return
	}
	fmt.Println(cost)
}
