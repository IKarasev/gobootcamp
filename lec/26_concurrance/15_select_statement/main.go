package main

import (
	"fmt"
	"strings"
	"time"
)

const lecHeader string = "Concurrency: Select Statement"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// select statement on channels allows
	// go routine operations to wait on multiple communication operations
	// select blocks until one of it cases can run
	// than it executes that case
	// if multiple cases are ready, they execute in random order

	// create two chan of string
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "Live long "
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "and prosper!"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			pf("1-Recieved %q\n", msg1)
		case msg2 := <-ch2:
			pf("2-Recieved %q\n", msg2)
		}
	}

}
