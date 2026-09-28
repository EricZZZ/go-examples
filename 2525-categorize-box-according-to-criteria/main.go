package main

import "fmt"

func categorizeBox(length int, width int, height int, mass int) string {
	v := length * width * height
	b := false
	h := false
	if mass >= 100 {
		h = true
	}
	if length >= 10000 || width >= 10000 || height >= 10000 || v >= 1000000000 {
		b = true
	}
	if b && h {
		return "Both"
	}
	if b {
		return "Bulky"
	}
	if h {
		return "Heavy"
	}
	return "Neither"
}

func main() {
	fmt.Println(categorizeBox(200, 50, 800, 50))
}
