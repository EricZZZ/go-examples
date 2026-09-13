package main

import "fmt"

func checkRecord(s string) bool {
	lc, ac, m := 0, 0, 0
	for _, c := range s {
		if c == int32('A') {
			ac++
		}
		if c == int32('L') {
			lc++
		} else {
			lc = 0
		}
		m = max(lc, m)
	}
	return ac < 2 && m < 3
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
}
