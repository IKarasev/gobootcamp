package main

import (
	"fmt"
	"sync"
)

func sum(f1, f2 float64, wg *sync.WaitGroup) {
	fmt.Printf("%.2f + %.2f = %.2f\n", f1, f2, f1+f2)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go sum(1.123, 2.345, &wg)
	go sum(5.83, 45.3, &wg)
	go sum(99.99, 2., &wg)

	wg.Wait()
}
