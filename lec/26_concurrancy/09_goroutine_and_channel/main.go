package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Concurrency: GoRoutine and Channels"

// count factorial using goroutine
func factorial(n int, ch chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	ch <- f // send f to channel
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	ch := make(chan int)
	defer close(ch) // close channel on main exits

	go factorial(10, ch)
	// main will wait till factorial() routine ends

	f := <-ch // get val from channel
	pln("factorial(10) =", f)

	pln(strings.Repeat("-", 10))

	// if more routines started
	for i := 1; i < 11; i++ {
		go factorial(i, ch)
		pf("factorial(%d) = %d\n", i, <-ch)
	}

}
