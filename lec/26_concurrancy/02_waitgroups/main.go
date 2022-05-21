package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const lecHeader string = "Concurrency: WaitGroups"

// functions to run in gorouting
// to use WaitGroups, the pointer to waitgroup var should passed
// to the function
func f1(wg *sync.WaitGroup) {
	fmt.Println("> f1() -- start")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i =", i)
		time.Sleep(time.Second)
	}

	fmt.Println("> f1() -- end")

	// in the end of concurrent function notify waitgroup that func has finished
	wg.Done()
}

func f2() {
	fmt.Println("> f2() -- start")

	for i := 0; i < 30; i += 10 {
		fmt.Println("f2, i =", i)
	}

	fmt.Println("> f2() -- end")
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// to ensure all gourinte to complite or sync them
	// WaitGroups are used

	// init wait group
	var wg sync.WaitGroup

	// before launching gourountines, define number of routines to waite for
	wg.Add(1)

	// launching gorouting
	// pass wait group pointer to func
	go f1(&wg)
	pln("No. of GoRoutines after go f1():", runtime.NumGoroutine())

	f2()

	// wait before goroutine has finished
	wg.Wait()

	pln("== End of main() ==")

}
