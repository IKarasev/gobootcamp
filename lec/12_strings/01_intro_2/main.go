package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Strings intro 2!")

	// rune type == int32
	v1, v2 := 'a', 'b'

	fmt.Printf("v1: %T = %v\n", v1, v1)
	fmt.Printf("v2: %T = %v\n", v2, v2)

	// non ascii chars
	str := "ţară"
	fmt.Println(str)
	fmt.Println(len(str)) // => 6, number of bytes, not chars
	fmt.Printf("Byte (note rune) at 1: %v\n", str[1])

	fmt.Println(strings.Repeat("-", 10))

	// iter str byte by byte
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d : %v : %c\n", i, str[i], str[i])
	}

	fmt.Println(strings.Repeat("-", 10))

	// iter str by runes
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:]) // r - rune, size - size of rune in bytes
		fmt.Printf("%d(%d) : %c\n", i, size, r)
		i += size // increase i by rune size in bytes
	}

	fmt.Println(strings.Repeat("-", 10))

	// decode runes auto
	for i, r := range str {
		fmt.Printf("%d: %c\n", i, r)
	}

}
