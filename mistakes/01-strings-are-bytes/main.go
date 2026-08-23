package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("len(hello):", len("hello"))
	fmt.Println("len(你好):", len("你好"))

	fmt.Println("\nlen(hello):", utf8.RuneCountInString("hello"))
	fmt.Println("len(你好):", utf8.RuneCountInString("你好"))
}
