package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func(f float64, wg *sync.WaitGroup) {
		fmt.Printf("sqrt of %f = %.4f\n", f, math.Sqrt(f))
		wg.Done()
	}(4.5846584, &wg)

	wg.Wait()
}
