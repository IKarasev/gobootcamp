package main

import (
	"fmt"
	"strings"
)

func main() {
	pln := fmt.Println
	pf := fmt.Printf

	pln("String package!")
	pln(strings.Repeat("-", 10))

	// contains substring
	r := strings.Contains("Somebody I used to know", "used")
	pln(r) // => true

	// contains any of given chars
	r = strings.ContainsAny("Victoria", "vx")
	pln(r) // => false, non of v or x are in Victoria (case sensitive)
	r = strings.ContainsAny("Victoria", "ix")
	pln(r) // => true

	// containes rune
	r = strings.ContainsRune("gaga", 'g')
	pln(r) // => true

	// occurances of substring (returnes number of occurrences)
	n := strings.Count("мама мыла раму, раму мыла мама", "мама")
	pln(n) // => 2
	// it can be used to count runes in string
	n = strings.Count("Four", "")
	pln(n) // => 5, number of runes + 1

	// lower upeer cases
	pln(strings.ToLower("SoMe STrinG in LOWER cAsE"))
	pln(strings.ToUpper("SoMe STrinG in upper cAsE"))

	// comparing strings
	pln("go" == "go") // => true
	pln("go" == "Go") // => false
	// case insensitive
	pln(strings.EqualFold("GO", "go")) // => true

	// replace char in str
	str := strings.Replace("192.168.0.1", ".", ":", -1) // -1 means replace all, other num means replace first n chars
	pln(str)
	str = strings.ReplaceAll("192.168.1.1", ".", ":")
	pln(str)

	// split string
	s := strings.Split("192.168.1.1", ".") // => return slice of strings splitted by separator
	pf("s (%T) = %#v\n", s, s)
	// splite rune by rune
	s = strings.Split("a b c,de", "")
	pf("s (%T) = %#v\n", s, s)

	// slice of strings into string
	s = []string{"I", "am", "learning", "GO", "!"}
	pln(strings.Join(s, " "))

	pln(strings.Repeat("-", 10))

	// split str by spaces & new lines
	str = "Some sort of \n apple"
	pln(str)
	fields := strings.Fields(str)
	pf("s (%T) = %#v\n", fields, fields)

	pln(strings.Repeat("-", 10))

	// remove spaces at start and end
	pf("%q\n", strings.TrimSpace("\t\n  Brew beer    \n   \t  "))

	// remove chars in str at start and end
	s1 := strings.Trim("....Brew beer! not bombs!!!!!??", ".!?")
	pf("%q\n", s1)

}
