package main

/*
The program will go throug list of URLs
Checks them, prints status
In this example it will be done using channels, goroutine
and anonymous function
*/

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const lecHeader string = "Concurrency: URL Serializer and Page Download with Channels\nand Ananymous func"

// 2. add channel param
// and put print into that channel
func CheckURL(url string, ch chan string) {
	resp, err := http.Get(url) // get request to url

	result := fmt.Sprintln("> URL:", url)

	if err != nil {
		result += fmt.Sprintln("is DOWN:", err)
	} else {
		//defer resp.Body.Close() // if no error close connection of func exit
		result += fmt.Sprintln("Status:", resp.StatusCode)
	}

	// send string result to channel
	fmt.Println(result)
	ch <- url
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	urls := []string{
		"https://golang.org",
		"https://www.google.com/aabb.html",
		"http://sssmmm.com",
	}

	// 1. create channel
	ch := make(chan string)

	for _, url := range urls {
		go CheckURL(url, ch)
	}

	// continiously check urls
	// for {
	// 	go CheckURL(<-ch, ch)
	// 	pln(strings.Repeat("=", 20))
	// 	time.Sleep(time.Second * 2) // to check every two seconds
	// }

	// another way:
	// for url := range ch {
	// 	time.Sleep(time.Second * 2)
	// 	go CheckURL(url, ch)
	// }

	// same using anonym func
	for url := range ch {
		go func(s string) {
			time.Sleep(time.Second * 2)
			go CheckURL(s, ch)
		}(url)
	}
}
