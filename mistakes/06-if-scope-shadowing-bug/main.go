package main

import (
	"errors"
	"fmt"
)

func pingDB() error { return errors.New("connection refused") }

func connectWrong(retry bool) error {
	var err error
	if retry {
		err := pingDB() // shadows the outer err variable
		fmt.Println(" inner err:", err)
	}
	return err
}

func main() {
	err := connectWrong(true)
	fmt.Println("outer err:", err)
}
