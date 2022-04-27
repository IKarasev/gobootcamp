package main

import "fmt"

func main() {
	var cities [2]string
	fmt.Printf("#1: %#v\n", cities)

	grades := [3]float64{1.1, 2.2}
	fmt.Printf("#2: %#v\n", grades)

	taskDone := [...]bool{true, false}
	_ = taskDone

	for i := 0; i < len(cities); i++ {
		fmt.Printf("city %d: %s\n", i, cities[i])
	}

	for i, v := range grades {
		fmt.Printf("grade %d: %.1f\n", i, v)
	}
}
