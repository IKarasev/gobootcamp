package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const lecHeader string = "Files: scan user input in stdin"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// create scanner
	scanner := bufio.NewScanner(os.Stdin)
	pf("scnnaer type: %T\n", scanner)

	// start user input
	pf("Write something: ")
	scanner.Scan()

	text := scanner.Text()   // text input result
	bytes := scanner.Bytes() // bytes input result
	pf("You entered:\n	text: %s\n	bytes: %v\n", text, bytes)

	// we can put file as stdin by
	// go run main.go < file
	// note that new line is an escape char

	// read input until exit code
	pln("-- Enter text (exit code: exit):")
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == "exit" {
			pln("-- Exited input --")
			break
		}
		pln("You entered:", scanner.Text())
	}

	// error handling
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
