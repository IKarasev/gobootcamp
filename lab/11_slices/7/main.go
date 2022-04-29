package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	cp := []string{}
	cp = append(cp, friends...)

	fmt.Printf("friends: %v\n", friends)
	fmt.Printf("copy: %v\n", cp)

	friends[0] = "Tony"
	cp[1] = "Ezekiel"

	fmt.Printf("friends: %v\n", friends)
	fmt.Printf("copy: %v\n", cp)
}
