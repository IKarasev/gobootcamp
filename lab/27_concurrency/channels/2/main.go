package main

import (
	"fmt"
	"os"
)

func main() {
	argNum := len(os.Args) - 1

	ch := make(chan string)

	if argNum > 0 {
		go func(ch chan string) {
			ch <- fmt.Sprintln("Recieved", argNum, "args")
		}(ch)

		fmt.Println(<-ch)
	} else {
		fmt.Println("No args provided")
	}
}
