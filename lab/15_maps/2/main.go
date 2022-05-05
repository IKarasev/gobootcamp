package main

import "fmt"

func main() {
	//var m1 map[int]bool
	//m1[5] = true
	// not init map
	m1 := map[int]bool{}
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}
	_, _ = m2, m3
	//fmt.Println(m2 == m3)
	// compare only to nill
	isequal := true
	for k, v := range m2 {
		if m3[k] != v {
			isequal = false
		}
	}
	fmt.Println(isequal)
}
