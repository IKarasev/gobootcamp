package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	const goNum = 50

	var wg sync.WaitGroup

	wg.Add(goNum)

	for i := 100.; i < 150.; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			fmt.Printf("sqrt of %.0f = %.4f\n", f, math.Sqrt(f))
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
