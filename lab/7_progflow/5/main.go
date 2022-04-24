package main

import (
	f "fmt"
	"time"
)

func main() {
	f.Println("Years from my birthday")
	birth := 1990
	cur := time.Now().Year()
	for i := birth; i <= cur; {
		f.Println(i)
		i++
	}
}
