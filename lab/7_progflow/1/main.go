package main

import f "fmt"

func main() {
	f.Println("[1;50] divisible by 7:")
	for i := 1; i < 51; i++ {
		if i%7 == 0 {
			f.Println(i)
		}
	}
}
