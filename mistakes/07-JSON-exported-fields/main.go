package main

import (
	"encoding/json"
	"fmt"
)

type UserWrong struct {
	Name  string
	email string
	Age   int
}

type UserRight struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	w := UserWrong{Name: "Eric", email: "example@example.com", Age: 44}
	out, _ := json.Marshal(w)
	fmt.Println("wrong:", string(out))

	r := UserRight{Name: "Eric", Email: "example@example.com", Age: 44}
	out, _ = json.Marshal(r)
	fmt.Println("right:", string(out))
}
