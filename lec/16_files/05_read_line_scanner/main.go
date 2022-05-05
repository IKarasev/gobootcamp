package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const lecHeader string = "Files: Read line by line with scanner!"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	file, err := os.Open("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	// init scanner
	scanner := bufio.NewScanner(file)

	success := scanner.Scan() // => bool, if file could be scanned
	if !success {
		err := scanner.Err() // if success false get error
		if err == nil {
			pln("Scan complited succesfully. Reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// get string text
	pln("First line: ", scanner.Text())
	// get bytes
	pln("First line in bytes:", scanner.Bytes())

	// scan the rest of text in file
	for scanner.Scan() {
		pln(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // check for errors in scan
		log.Fatal(err)
	}

	file.Close()

	pln(strings.Repeat("-", 10))

	// scan word by word
	file, err = os.Open("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		pln(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // check for errors in scan
		log.Fatal(err)
	}

	file.Close()

	pln(strings.Repeat("-", 10))

	// scan by rune
	file, err = os.Open("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		pln(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // check for errors in scan
		log.Fatal(err)
	}

	file.Close()
}
