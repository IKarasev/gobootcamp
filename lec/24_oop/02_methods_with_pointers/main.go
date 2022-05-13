package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "OOP: methods with pointer params"

// define new struct
type car struct {
	brand string
	model string
	price int
}

// creating a method for struct car that changes fields
func (c *car) setPrice(newPrice int) {
	// pointer as method reciever allows to change it fields directrly
	fmt.Printf("> car.setPrice(%d)\n", newPrice)
	(*c).price = newPrice
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// create car var
	myCar := car{
		brand: "Audi",
		model: "A6",
		price: 40000,
	}
	pln("myCar =", myCar)

	// change car price with method
	myCar.setPrice(55000)
	pln("myCar =", myCar)

	// call method from pointer
	ptr := &myCar
	pln("ptr =", ptr)

	ptr.setPrice(60000) // setPrice has pointer as reciever, so can call it directly from pointer to var var
	pln("myCar =", myCar)
}
