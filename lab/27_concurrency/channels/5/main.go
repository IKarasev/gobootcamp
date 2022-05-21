package main

import "fmt"

func power(n int, c chan int) {
	c <- n * n
}

func main() {
	const goNum = 50

	ch := make(chan int)

	for i := 1; i <= goNum; i++ {
		go func(n int) {
			ch <- n * n
		}(i)
	}

	for i := 1; i <= goNum; i++ {
		fmt.Println(<-ch)
	}

}
