package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Structs: anonymous structs and fields"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// useful when struct type is not reused
	// def anonymous struct
	person := struct {
		// def fields
		firstName, lastName string
		age                 int
	}{
		// def values
		firstName: "Antony",
		lastName:  "Bridge",
		age:       25,
	}
	pf("%#v\n", person)
	pf("%+v\n", person) // field names with vals
	pf("%-v\n", person) // field vals

	pln(strings.Repeat("-", 10))

	// anonymous fields
	type Book struct {
		string
		int
		bool
	}

	b1 := Book{`"The Island of Dr. Moreau" by Herbert Wells`, 50, true}
	pf("b1 = %#v\n", b1)
	pln("b1.string =", b1.string)
	// field types converted to field names

	// can mix named and unnamed fields
	type Empl struct {
		name   string
		salary int
		bool
	}

	// john := Empl{
	// 	name: "John Malkovich",
	// 	salary: 10000,
	// 	true,
	// } // ERROR: can't mix named and unnamed field in init struct val

	john := Empl{"John Malkovich", 10000, true}
	pf("john = %#v\n", john)
	// unnamed field refer by type name
	john.bool = false
	pf("john = %#v\n", john)

}
