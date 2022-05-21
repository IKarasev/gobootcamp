package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const lecHeader string = "Concurrency: Spawning GoRoutines"

// functions to run in gorouting
func f1() {
	fmt.Println("> f1() -- start")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i =", i)
	}

	fmt.Println("> f1() -- end")
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

	pln("No. of CPUs: ", runtime.NumCPU())
	pln("No. of GoRoutines:", runtime.NumGoroutine())
	pln("OS:", runtime.GOOS)
	pln("Arch:", runtime.GOARCH)

	// how many OS threads can be done in parallel on this machine
	pln("GOMAXPROX:", runtime.GOMAXPROCS(0)) // 0 - show current setting, by default == # of CPUs
	// other value will set max procs number

	// launching goroutine
	go f1()
	pln("No. of GoRoutines after go f1():", runtime.NumGoroutine())

	f2()

	// main() finishes before f1()
	// set sleep time for main()
	time.Sleep(time.Second * 2)

	pln("== End of main() ==")

}
