package main

import "fmt"

func main() {
	s := []string{"one", "two", "three"}
	for i, v := range s {
		fmt.Println("index", i, "value", v)
	}
}
