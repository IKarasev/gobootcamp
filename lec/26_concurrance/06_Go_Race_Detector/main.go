package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const lecHeader string = "Concurrency: Data Racing"

// here we'll try to detect race error
// for that "go run -race" is used on running program

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	const gr = 300 //num of routines

	var wg sync.WaitGroup // wait group
	wg.Add(gr * 2)

	var n int = 0 // var to work with for routines

	// starting go routines
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}
	wg.Wait()

	pln("n =", n) // not always == 0
	// that is what called DataRace
	// during execution of routines, at some point several routines try to access n var at same time
	// as a result only one of them complited others passes, as they can't modify n
}
