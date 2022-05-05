package main

import (
	"fmt"
	"strings"
)

const lecHeader string = "Maps intro!"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// map header
	// on init, go creates a pointer to map header in memory
	// pairs stored in memory, and pointer to it start in map header

	friends := map[string]int{"Dan": 12, "Marry": 16, "John": 10}
	pln("friends:", friends)

	neib := friends
	pln("neib:", neib)

	pln(strings.Repeat("-", 10))

	// mod friends affect neib, as map is not copied. Copied only header
	friends["Dan"] = 22
	pln(`> friends["Dan"] = 22`)
	pln("friends:", friends)
	pln("neib:", neib)

	// vice versa with neib
	neib["Marry"] = 25
	pln(`> neib["Marry"] = 25`)
	pln("friends:", friends)
	pln("neib:", neib)

	// and del
	delete(friends, "John")
	pln(`> delete(friends, "John")`)
	pln("friends:", friends)
	pln("neib:", neib)

	pln(strings.Repeat("-", 10))

	// to copy map - init new, and init via for loop
	fr := map[string]int{}
	for k, v := range friends {
		fr[k] = v
	}
	pln("friends:", friends)
	pln("fr:", fr)

	fr["Dan"] = 40
	pln(`> fr["Dan"] = 40`)
	pln("friends:", friends)
	pln("fr:", fr)
}
