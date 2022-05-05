package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const lecHeader string = "Files: read in bytes!"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// read n bytes from file
	bSlice := make([]byte, 2) // slice size sets # of bytes to read from file

	bRead, err := io.ReadFull(file, bSlice) // error if # of bytes in file < slice size
	if err != nil {
		log.Fatal(err)
	}

	pln("Read", bRead, "bytes")
	pf("Data read: %s\n", bSlice)

	pln(strings.Repeat("-", 10))

	// read all from file
	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	pln(data)
	pf("%s\n", data)

	pln(strings.Repeat("-", 10))

	// read and close
	data, err = ioutil.ReadFile("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	pf("%s\n", data)

}
