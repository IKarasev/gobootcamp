package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Slicing strings!")
	s1 := "Somewhere over the rainbow"
	fmt.Println("s1 =", s1)
	fmt.Println("s1[2:5] =", s1[2:5]) // => mew

	s2 := "Сила в правде"
	fmt.Println("s2 =", s2)
	fmt.Println("s2[:4] =", s2[:4]) // => Си , slice by bytes, not runes

	fmt.Println(strings.Repeat("-", 10))

	// slicing string by runes (not bytes):
	rs := []rune(s2) // conver str to slice of runes
	fmt.Printf("%#v\n", rs)

	// print first 4 letters
	fmt.Println(string(rs[0:4])) // => Сила

	// NOTE: Converting str to rune and back is not efficient, as new backing array is created each time
}
