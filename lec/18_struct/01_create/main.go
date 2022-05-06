package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Structs: creating!"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// init
	type book struct {
		title  string
		author string
		year   int
	}

	// if some fields of same type we can do:
	type bookNew struct {
		title, author string
		year          int
	}

	// create var of struct
	myBook := book{"Hyperion", "Dan Simmons", 1989}
	pln(myBook)

	// named init var
	myBook2 := book{
		title:  "A Farewell to Arms",
		author: "Ernest Miller Hemingway",
		year:   1929,
	}
	pln(myBook2)

	// only one field can be specified
	myBook3 := book{title: "The Book of Jungle"}
	pf("%#v\n", myBook3)

}
