package main

import "fmt"

func main() {
	fmt.Println("Keyed array!")
	grades := [3]int{
		2: 5,
		0: 4,
		1: 7,
	}
	fmt.Println(grades)

	acc := [3]int{2: 15}
	fmt.Println(acc)

	// with auto len
	ppl := [...]string{
		2: "Ilya",
		5: "John",
		3: "Jane",
		8: "Marry",
	}
	fmt.Printf("%#v\n", ppl)
	fmt.Printf("Len of ppl: %d\n", len(ppl))

	nums := [...]int{
		6:  1,
		11, // will be put as last elem after bigger key, => index 7
		2:  3,
	}
	fmt.Printf("%#v\n", nums)

}
