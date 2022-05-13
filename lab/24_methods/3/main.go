package main

import "fmt"

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	(*b).price *= 0.9
}

func main() {
	myBook := book{
		title: "Dune",
		price: 100,
	}
	fmt.Printf("myBook: %+v\n", myBook)
	fmt.Println("vat =", myBook.vat())

	myBook.discount()
	fmt.Printf("myBook: %+v\n", myBook)
}
