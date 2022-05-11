package main

import (
	"fmt"
	"log"
	"strconv"
)

func myFunc(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	nn, _ := strconv.Atoi(s + s)
	nnn, _ := strconv.Atoi(s + s + s)

	res := n + nn + nnn

	return res

}

func main() {
	fmt.Println("myFinc(5) =", myFunc("5"))
	fmt.Println("myFinc(9) =", myFunc("9"))
	fmt.Println("myFinc(sss) =", myFunc("sss"))
}
