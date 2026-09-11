package main

import "fmt"

func totalNumbers(digits []int) int {
	set := make(map[int]struct{})

	for i, _ := range digits {
		if digits[i] == 0 {
			continue
		}
		for j, _ := range digits {
			if i == j {
				continue
			}

			for k, _ := range digits {
				if k == i || k == j {
					continue
				}

				if digits[k]%2 != 0 {
					continue
				}

				num := digits[i]*100 + digits[j]*10 + digits[k]
				set[num] = struct{}{}
			}
		}
	}
	return len(set)
}

func main() {
	fmt.Println(totalNumbers([]int{0, 2, 2}))
}
