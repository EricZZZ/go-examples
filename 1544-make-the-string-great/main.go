package main

import "fmt"

func makeGood(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if len(stack) > 0 && abs(int(stack[len(stack)-1])-int(c)) == 32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
}
