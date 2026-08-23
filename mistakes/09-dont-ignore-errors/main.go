package main

import (
	"fmt"
	"os"
)

func loadWrong(path string) int64 {
	info, _ := os.Stat(path)
	return info.Size() // panic if the file does not exist
}

func loadRight(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("stat: %w", err)
	}
	return info.Size(), nil
}

func main() {
	size, err := loadRight("missing.json")
	if err != nil {
		fmt.Println("right:", err)
	} else {
		fmt.Println("right:", size)
	}

	fmt.Println("wrong:", loadWrong("missing.json"))
}
