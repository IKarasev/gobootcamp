package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Structs: fields"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	type book struct {
		title  string
		author string
		year   int
	}

	myBook := book{title: "At the Mountains of Madness"}
	pf("%#v\n", myBook)
	pln("title:", myBook.title)
	// pln("pages:", myBook.pages) // ERROR: myBook.pages undefined (type book has no field or method pages)
	// fields can't be changed\deleted\added at a runtime

	// change field val
	myBook.author = "Howard Lovecraft"
	myBook.year = 1939
	pf("%+v\n", myBook) // print fields with vals only

	pln(strings.Repeat("-", 10))

	// can compare structs directrly
	myBook2 := book{
		title:  "Some Book", // different title
		author: "Howard Lovecraft",
		year:   1939,
	}

	pln("myBook == myBook2:", myBook == myBook2)

	myBook2.title = "At the Mountains of Madness"
	pln("myBook == myBook2:", myBook == myBook2)

	pln(strings.Repeat("-", 10))

	// copying structs
	myBook3 := myBook
	pf("%#v\n", myBook3)

	// change copy does not affect origin
	myBook3.title = "Some new title"
	pf("%#v\n", myBook3)
	pf("%#v\n", myBook)

}
