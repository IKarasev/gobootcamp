package main

import "fmt"

// init func called before main
func init() {
	fmt.Println("=> Init func")
}

// we can have multiple init funcs
// they be called in order as in the code
func init() {
	fmt.Println("=> Second Init func")
}

// init funcs are used to do smth before main() called
// eg: to define global vars, if it can't be done other way

// declare global array var
var nums [10]int // if we print this array withougt init it will be = [0 0 0 0 0 0 0 0 0 0]

// if we need to init array before main() we ca do it in init
func init() {
	for i := 0; i < 10; i++ {
		nums[i] = i + 1
	}
}

// now if we print nums => [1 2 3 4 5 6 7 8 9 10]

func main() {
	fmt.Println("=> Main func")

	// init() // => error!
	// we can't call init() explicitly

	fmt.Println(nums)
}
