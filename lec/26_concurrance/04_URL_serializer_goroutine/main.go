package main

/*
The program will go throug list of URLs
Checks them, prints status
if webpage on url up - saves responce body to a file
now we do it using goroutine
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

const lecHeader string = "Concurrency: URL Serializer and Page Download with GoRoutines"

// adding pointer to waitGroup
func CheckAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url) // get request to url

	fmt.Println("=> URL:", url)

	if err != nil {
		fmt.Println("is DOWN:", err)
	} else {
		defer resp.Body.Close() // if no error close connection of func exit
		fmt.Println("Status:", resp.StatusCode)

		// if status ok - save body to file
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			fileName := strings.Split(url, "//")[1]
			fileName += ".html"

			fmt.Printf("   Write body to %s\n", fileName)

			err = ioutil.WriteFile(fileName, bodyBytes, 0664) // write to file

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// don't forget to complite task in WaitGroup
	wg.Done()
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

	// create waitgroup
	var wg sync.WaitGroup

	// number of goroutines to wait
	wg.Add(len(urls))

	for i, url := range urls {
		// launch checks as goroutine
		// don't forget pointer to wg
		go CheckAndSaveBody(url, &wg)
		pf("== [%d] check %q ==\n", i, url)
	}

	// make program gait for goroutine to finish
	wg.Wait()
}
