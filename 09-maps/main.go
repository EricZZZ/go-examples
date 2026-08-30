package main

import (
	"fmt"
	"maps"
)

func main() {
	inventory := map[string]int{
		"pens":    10,
		"pencils": 5,
		"erasers": 3,
	}

	inventory["notebook"] = 12

	v, ok := inventory["paper clips"]

	if ok {
		fmt.Println("The inventory has ", v, " paper clips")
	} else {
		fmt.Println("Paper clips are not in the inventory")
	}

	fmt.Println(inventory)

	// delete map
	delete(inventory, "pens")

	fmt.Println(inventory)

	for item, quantity := range inventory {
		fmt.Printf("There are %d %s in the inventory\n", quantity, item)
	}

	inventory1 := map[string]int{
		"pens":    10,
		"pencils": 5,
		"erasers": 3,
	}

	inventory2 := map[string]int{
		"pens":    10,
		"pencils": 5,
		"erasers": 3,
	}

	// compare maps
	if maps.Equal(inventory1, inventory2) {
		fmt.Println("The maps are equal")
	} else {
		fmt.Println("The maps are not equal")
	}

	// clear map
	fmt.Println("clear before:", inventory1)
	clear(inventory1)
	fmt.Println("clear after::", inventory1)
}
