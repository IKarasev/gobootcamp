package main

import "fmt"

type Person struct {
	name   string
	age    int
	colors []string
}

func main() {
	me := Person{
		name:   "My name",
		age:    30,
		colors: []string{"red", "yellow", "green", "black"},
	}

	for i, v := range me.colors {
		fmt.Println(i, ":", v)
	}
}
