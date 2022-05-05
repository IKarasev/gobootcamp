package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const lecHeader string = "Files: write to file"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	file, err := os.OpenFile(
		"a.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // closes file when exits current function (in this case main)

	// to write string in bytes we need to convert str to byte slice

	bSlice := []byte("Let learn GO!")
	bWritten, err := file.Write(bSlice) // returns # of bytes written and error

	if err != nil {
		log.Fatal(err)
	}
	pln("Bytes written:", bWritten)

	// another way with ioutil
	bs := []byte("It's not that hard!")
	err = ioutil.WriteFile("b.txt", bs, 0644) // if file not exist - creates it, otherwise truncates it
	// input:
	//	file name / path
	//	bytes to write
	//	pernissions

	if err != nil {
		log.Fatal(err)
	}
}
