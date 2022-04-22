package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Simple if!")
	i, err := strconv.Atoi("44")

	if err != nil {
		fmt.Println("Error converting str to int")
	} else {
		fmt.Println("i = ", i)
	}

	// short if
	if i, err := strconv.Atoi("20"); err == nil {
		fmt.Println("i = ", i)
	} else {
		fmt.Println("Error converting str to int")
	}

	if args := os.Args; len(args) < 2 {
		fmt.Println("At least one arg is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("First arg must int")
	} else {
		fmt.Printf("%d km = %v miles\n", km, float64(km)*0.629)
	}

}
