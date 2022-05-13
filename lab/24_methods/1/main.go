package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func main() {
	var myMoney money = 20.365
	myMoney.print()
}
