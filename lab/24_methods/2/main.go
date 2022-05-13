package main

import "fmt"

type money float64

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func main() {
	var myMoney money = 20.365

	s := myMoney.printStr()
	fmt.Printf("%T\n", s)
	fmt.Println(s)
}
