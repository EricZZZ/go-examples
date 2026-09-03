package main

func inventoryManagement(stock []int) int {
	l := len(stock)
	t := make(map[int]int, l)
	for _, v := range stock {
		t[v]++
		if t[v] > l/2 {
			return v
		}
	}
	return 0
}

func main() {
	stock := []int{-1, 1, 1, 1, 2, 1}
	count := inventoryManagement(stock)
	println(count)
}
