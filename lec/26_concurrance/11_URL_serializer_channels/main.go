package main

/*
The program will go throug list of URLs
Checks them, prints status
if webpage on url up - saves responce body to a file
in this example it will be done using channels and goroutine
*/

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const lecHeader string = "Concurrency: URL Serializer and Page Download with Channels"

// 2. add channel param
// and put print into that channel
func CheckAndSaveBody(url string, ch chan string) {
	resp, err := http.Get(url) // get request to url

	result := fmt.Sprintln("> URL:", url)

	if err != nil {
		result += fmt.Sprintln("is DOWN:", err)
	} else {
		defer resp.Body.Close() // if no error close connection of func exit
		result += fmt.Sprintln("Status:", resp.StatusCode)

		// if status ok - save body to file
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			fileName := strings.Split(url, "//")[1]
			fileName += ".html"

			result += fmt.Sprintln("Write body to", fileName)

			err = ioutil.WriteFile(fileName, bodyBytes, 0664) // write to file

			if err != nil {
				// log.Fatal(err) als change to text output
				result += fmt.Sprintln("ERROR writing to file:", err)
			}
		}
	}

	// send string result to channel
	ch <- result
}

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://www.google.com/aabb.html",
		"https://ya.ru",
		"http://sssmmm.com",
	}

	// 1. create channel
	ch := make(chan string)

	for _, url := range urls {
		go CheckAndSaveBody(url, ch)
	}

	// print vals from channel
	for i := 0; i < len(urls); i++ {
		pln(<-ch)
	}
}
