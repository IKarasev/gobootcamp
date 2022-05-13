package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const lecHeader string = "Functions: DEFER"

func foo() {
	fmt.Println("> This is foo()")
}

func bar() {
	fmt.Println("> This is bar()")
}

func foobar() {
	fmt.Println("> This is foobar()")
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	defer foo() // foo() will be called on main() end
	bar()

	pln("Some code after foo and bar")

	defer foobar() // defer funcs will be executed in reverse order it was defined
	// so, after main() ends foobar() called first, and foo() second

	pln("Some code after foobar")

	//defer can be used to ensure cleanup after func ends, eg with files
	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pln("Code to work with file")
	pln("File will be closed after main() end")
}
