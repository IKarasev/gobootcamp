package main

import "fmt"

func main() {
	fmt.Println("Strings intro 1!")

	s1 := "hello, go"
	fmt.Println(s1)

	// quotes in string
	fmt.Println("James Cameron: \"If you set your goals ridiculously high and it's a failure, you will fail above everyone else's success.\"")
	fmt.Println(`Benjamin Franklin: "Tell me and I forget. Teach me and I remember. Involve me and I learn."`)

	// raw srting
	s2 := `Learng Go \n and go!`
	fmt.Println(s2)

	// same new line
	fmt.Println("Brand: Asics\nPrice: 100")
	fmt.Println(`Brand: Asics
Price: 100`)

	// concat strings
	s3 := "Best " + "day " + "ever"
	fmt.Println(s3 + "!")

	// indexing
	fmt.Println("Char at 0:", s3[0]) // => 66

	// string is slice of bytes
	// change letter in index

	// s3[3] = 'T' // error, need byte ascii code of letter

}
