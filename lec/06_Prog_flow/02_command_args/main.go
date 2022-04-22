package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Com lines args!")
	fmt.Println("given args:", len(os.Args))
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("arg 1: ", os.Args[1])
	fmt.Println("arg 2: ", os.Args[2])

	var res, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("first arg to float: %f\n", res)
	fmt.Printf("convert errors: %v\n", err)
}
