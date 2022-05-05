package main

import (
	"fmt"
	"strings"
)

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("Maps intro!\n")
	pln(strings.Repeat("-", 10))

	// map ~ dict in python
	// keys must be of same type
	// values must be of same type

	// decl
	var empl map[string]string // new map with str keys and str vals
	pf("%T = %v\n", empl, empl)
	pln("empl is empty (== nil):", empl == nil)
}
