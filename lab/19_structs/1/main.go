package main

import "fmt"

type Person struct {
	name   string
	age    int
	colors []string
}

func main() {
	me := Person{
		name: "My name",
		age:  30,
	}
	you := Person{
		name:   "Your name",
		age:    32,
		colors: []string{"green", "blue"},
	}

	fmt.Printf("me = %#v\n", me)
	fmt.Printf("%+v\n", me)
	fmt.Printf("you = %#v\n", you)
	fmt.Printf("%+v\n", you)
}
