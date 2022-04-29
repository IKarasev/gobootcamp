package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("At least 2 args are required")
	} else if len(os.Args) > 11 {
		fmt.Println("Max of 10 args are supported")
	} else {
		sum := 0.
		prod := 1.

		for i, v := range os.Args[1:] {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Printf("Arg №%d is not int (%v). It won't be counted\n", i, v)
				continue
			}
			sum = sum + num
			prod = prod * num
		}
		fmt.Println("sum =", sum)
		fmt.Println("product =", prod)
	}
}
