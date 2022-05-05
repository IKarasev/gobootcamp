package main

import (
	"fmt"
	"strings"
)

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("Maps: in work!\n")
	pln(strings.Repeat("-", 10))

	// decl
	var empl map[string]string // new map with str keys and str vals
	pf("%T = %#v\n", empl, empl)
	pln("empl is empty (== nil):", empl == nil)

	pf("Num of pairs: %d\n", len(empl))
	pf("val for %q: %q\n", "Antony", empl["Antony"]) // => empty str, as key doen't exist
	// if key doesn't exist map returnes zero val of val type

	var acc map[string]float64
	pf("val for %q: %.2f", "Moana", acc["Moana"])

	pln(strings.Repeat("-", 10))

	// slice cannot be key, but array can!
	// only comparable types could be keys
	//var m map[[]int]string // => invalid map key type []int
	var m map[[2]int]string
	pf("m:\n%T = %#v\n", m, m)

	pln(strings.Repeat("-", 10))

	// inserting pairs
	// empl["Antony"] = "Cleaner" // => panic: assignment to entry in nil map
	// can't insert key:val to uninit map

	// init map
	empl = map[string]string{}
	empl["Antony"] = "Cleaner"
	pf("empl:\n%T = %#v\n", empl, empl)

	ppl := map[string]float64{}
	ppl["TJ"] = 14.2
	ppl["Marry"] = 12.8
	pf("ppl:\n%T = %#v\n", ppl, ppl)

	// init and init empty map
	m1 := make(map[int]int)
	pln(m1)
	m1[4] = 44
	pln(m1)

	// another way
	blnc := map[string]float64{
		"USD": 334.2,
		"EUR": 1000,
		"CNY": 204.54,
	}
	pln(blnc)

	blnc["USD"] = 500.0
	blnc["GBP"] = 333.33
	pln(blnc)

	// val for non existing key is zero of val type
	pln("CHF:", blnc["CHF"])

	// how to distinguish betwing non-ex key and key with zero val
	v, ok := blnc["CHF"]        // => ok is false, as key doesn't exist
	pln("CHF:", v, "; ok:", ok) // => CHF: 0 ; ok: false

	// loop through
	// NOTE: order of key:val can be different from assignment
	pln(strings.Repeat("-", 10))

	for k, v := range blnc {
		pf("%s: %.2f\n", k, v)
	}

	// deleting keys
	delete(blnc, "GBP")
	pln(blnc)
}
