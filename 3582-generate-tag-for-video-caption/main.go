package main

import (
	"fmt"
	"strings"
	"unicode"
)

func generateTag(caption string) string {
	var b strings.Builder
	for _, ch := range caption {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == ' ' {
			b.WriteRune(ch)
		}
	}
	caption = strings.TrimSpace(b.String())
	strs := strings.Split(caption, " ")

	var builder strings.Builder
	builder.WriteString("#")
	for i, v := range strs {

		if v == "" {
			continue
		}
		if builder.Len() > 100 {
			break
		}
		v = strings.ToLower(v)
		if i != 0 {
			r := []rune(v)
			r[0] = unicode.ToUpper(r[0])
			v = string(r)
		}
		builder.WriteString(v[0:min(100-builder.Len(), len(v))])
	}
	return builder.String()
}

func main() {
	fmt.Println(generateTag("Leetcode daily streak achieved"))
}
