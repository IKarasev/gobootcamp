package main

/*
The program will go throug list of URLs
Checks them, prints status
if webpage on url up - saves responce body to a file
in this example it will be done sequentially
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const lecHeader string = "Concurrency: URL Serializer and Page Download"

func CheckAndSaveBody(url string) {
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

	for _, url := range urls {
		CheckAndSaveBody(url)
		pln(strings.Repeat("=", 20))
	}
}
