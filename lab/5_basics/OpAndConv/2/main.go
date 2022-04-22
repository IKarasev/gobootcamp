package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	i1 := strconv.Itoa(i)
	s2_2, err := strconv.Atoi(s2)
	f1 := fmt.Sprintf("%f", f)
	s1_2, err := strconv.ParseFloat(s1, 64)

	_ = err
	fmt.Printf("new i type %T, val %v\n", i1, i1)
	fmt.Printf("new s2 type %T, val %v\n", s2_2, s2_2)
	fmt.Printf("new f type %T, val %v\n", f1, f1)
	fmt.Printf("new s1 type %T, val %v\n", s1_2, s1_2)
}
