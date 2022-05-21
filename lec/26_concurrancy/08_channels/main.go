package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Concurrency: Channels example"

// create func for concurent operations with channel
func f1(n int, ch chan int) {
	// send n to channel
	ch <- n

}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// init channel
	ch := make(chan int) // by-directional channel

	// alsom 1-dyrection (uni-direction) channels supported
	chRecieve := make(<-chan string) // recieve only
	chSend := make(chan<- string)    // send only

	pf("ch ~ %T\n", ch)
	pf("chRecieve ~ %T\n", chRecieve)
	pf("chSend ~ %T\n", chSend)

	pln(strings.Repeat("-", 10))

	// run f1
	// main routine and f1() routine are connected
	go f1(69, ch)

	// get val from channel
	n := <-ch
	pln("ch val = ", n)

	pln("== Main exit ==")
}
