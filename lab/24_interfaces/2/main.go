package main

import "fmt"

type empty interface{}

func main() {
	var e empty

	fmt.Println(`> var e empty`)
	fmt.Printf("Type: %T\nValue: %#v\n\n", e, e)

	e = 25
	fmt.Println(`> e = 25`)
	fmt.Printf("Type: %T\nValue: %#v\n\n", e, e)

	e = 5.789
	fmt.Println(`> e = 5.789`)
	fmt.Printf("Type: %T\nValue: %#v\n\n", e, e)

	e = []int{1, 2, 3}
	fmt.Println(`> e = []int{1, 2, 3}`)
	fmt.Printf("Type: %T\nValue: %#v\n\n", e, e)

	e = append(e.([]int), 4)
	fmt.Println(`> e = append(e.([]int), 4)`)
	fmt.Printf("Type: %T\nValue: %#v\n\n", e, e)
}
