package main

import "fmt"

func countPrefixSuffixPairs(words []string) int {
	sum := 0
	for i := range len(words) - 1 {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				sum++
			}
		}
	}
	return sum
}

func isPrefixAndSuffix(str1 string, str2 string) bool {
	if len(str1) > len(str2) {
		return false
	}
	pre := str2[:len(str1)]
	if pre != str1 {
		return false
	}
	suf := str2[len(str2)-len(str1):]
	if suf != str1 {
		return false
	}
	return true
}

func main() {
	// fmt.Println(isPrefixAndSuffix("aba", "aa"))
	fmt.Println(countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa"}))
}
