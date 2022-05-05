package main

import (
	"fmt"
	"strings"
)

const lecHeader string = ""

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))
}
