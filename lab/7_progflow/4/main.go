package main

import f "fmt"

func main() {
	f.Println("[1;500] divisible by 7 and 5:")
	for i := 1; i < 501; i++ {
		if i%7 == 0 && i%5 == 0 {
			f.Println(i)
		}
	}
}
