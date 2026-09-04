package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "Eric", age: 18}
	component := hello(p)
	basic_syntax_component := headerTemplate("Good")
	button_component := button("Eric", "Say Hello")

	http.Handle("/", templ.Handler(component))
	http.Handle("/basic", templ.Handler(basic_syntax_component))
	http.Handle("/button", templ.Handler(button_component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
