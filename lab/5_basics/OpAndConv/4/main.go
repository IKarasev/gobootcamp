package main

import "fmt"

func main() {
	var StoE float64 = 149.6
	var speed int64 = 299792458
	time := (StoE * 1000000 * 1000) / float64(speed)

	fmt.Printf("Time = %.3f seconds", time)
}
