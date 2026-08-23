package main

import "fmt"

func main() {
	roles := []string{"viewer", "editor"}
	addRole(roles)
	fmt.Println(roles)
	addRole_1(&roles)
	fmt.Println(roles)
	roles = addRole_2(roles)
	fmt.Println(roles)
}

func addRole(roles []string) {
	roles = append(roles, "main")
}

func addRole_1(roles *[]string) {
	*roles = append(*roles, "admin")
}

func addRole_2(roles []string) []string {
	roles = append(roles, "admin")
	return roles
}
