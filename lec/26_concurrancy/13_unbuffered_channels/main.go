package main

import (
	"fmt"
	"strings"
	"time"
)

const lecHeader string = "Concurrency: Unbuffered Channels"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	ch1 := make(chan int) // unbuddered channel

	// second arg could be passed to make, named channel capacity
	ch2 := make(chan int, 3) // buddered channel
	_ = ch2

	// try it
	go func(c chan int) {
		pln("Before sendind to chan")
		c <- 10
		pln("After sendind to chan")
	}(ch1)

	pln("main() sleeps for 2 sec")
	time.Sleep(time.Second * 2)

	pln("main() recieve data from channel")
	d := <-ch1
	pln("recieved val:", d)

	time.Sleep(time.Second)
}
