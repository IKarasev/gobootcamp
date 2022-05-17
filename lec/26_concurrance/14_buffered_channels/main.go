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

	ch2 := make(chan int, 3) // buddered channel

	// try it
	go func(c chan int) {
		for i := 1; i < 6; i++ {
			pf("[routine %d] Before sendind to chan\n", i)
			c <- i
			pf("[routine %d] After sendind to chan\n", i)
		}
		close(c)
	}(ch2)

	pln("main() sleeps for 2 sec")
	time.Sleep(time.Second * 2)

	pln("main() recieve data from channel")
	for v := range ch2 {
		pln("data =", v)
	}

	// try get value from channel after all data is recieved
	pln("After all data recieved chan val =", <-ch2)

	// try send data to closed channel
	// ch2 <- 22 // => panic: send on closed channel
}
