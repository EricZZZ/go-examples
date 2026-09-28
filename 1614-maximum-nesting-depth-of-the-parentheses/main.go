package main

import "fmt"

func maxDepth(s string) int {
	l, r := 0, 0
	ans := 0
	for _, v := range s {
		if v == '(' {
			l++
		}
		if v == ')' {
			r++
			l--
		}
		ans = max(ans, l)
	}
	return ans
}

func main() {
	fmt.Println(maxDepth("()(())((()()))"))
}
