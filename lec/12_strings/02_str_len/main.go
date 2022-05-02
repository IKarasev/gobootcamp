package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Strings length in runes and bytes!")

	// string is sequence of bytes
	s1 := "Golang"
	fmt.Println("s1 =", s1)
	fmt.Println("len(s1) =", len(s1)) // => 6, ascii letters are 1 byte each => len in bytes and len of str are same

	name := "Ярослав"
	fmt.Println("name =", name)
	fmt.Println("len(name) =", len(name)) // => 14, non ascii chars

	fmt.Println("len(\"Ж\") =", len("Ж")) // => 2, 2 bytes for char Ж

	// count runes in str
	n := utf8.RuneCountInString(name)
	fmt.Println("name", name, "has", n, "chars")
}
