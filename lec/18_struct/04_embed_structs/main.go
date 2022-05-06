package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Structs: embedded structs!"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	type Contact struct {
		address, email string
		phone          int
	}

	type Employee struct {
		name     string
		salary   int
		contacts Contact
	}

	john := Employee{
		name:   "John Malkovich",
		salary: 5000,
		contacts: Contact{
			address: "NY",
			email:   "jm@gmail.com",
			phone:   123456789,
		},
	}

	pf("john = %#v\n", john)
	pf("%+v\n", john)

	pln("Salary raise!")
	john.salary = 8000
	pln("Relocation")
	john.contacts.address = "Budapest"
	pf("john = %#v\n", john)

	pln(strings.Repeat("-", 10))

	// contact var can be used separatly and assigned to othe struct embedded field
	newContact := Contact{
		address: "Adelaide",
		email:   "new_mail@somehost.com",
		phone:   111111,
	}

	mike := Employee{
		name:     "Mike Vazovski",
		salary:   3000,
		contacts: newContact,
	}

	pf("newContact = %+v\n", newContact)
	pf("mike = %+v\n", mike)

	pln(strings.Repeat("-", 10))

	// Change newContact, not affect mike
	pln("Change newContact, not affect mike")
	newContact.email = "other@mail.en"
	pf("newContact = %+v\n", newContact)
	pf("mike = %+v\n", mike)

}
