package main

import (
	"fmt"
	"strings"
)

func searchItem(sl []string, s string) bool {
	for _, v := range sl {
		if strings.EqualFold(v, s) {
			return true
		}
	}

	return false
}

func main() {
	sl := []string{"pig", "dog", "tiger", "chair"}
	fmt.Println("dog in slice:", searchItem(sl, "dog"))
	fmt.Println("buddy in slice:", searchItem(sl, "buddy"))
}
