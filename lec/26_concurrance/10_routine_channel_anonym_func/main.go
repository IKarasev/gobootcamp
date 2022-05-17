package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Concurrency: GoRoutine, Channels and Anonymous Function"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	ch := make(chan int)
	defer close(ch) // close channel on main exits

	// span routines with anonym func
	for i := 1; i < 11; i++ {
		go func(n int, ch chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}
			ch <- f // send f to channel
		}(i, ch)

		pf("factorial(%d) = %d\n", i, <-ch)
	}

}
