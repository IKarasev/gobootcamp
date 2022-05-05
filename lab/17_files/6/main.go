package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fileName := "info.txt"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text := []byte("The Go gopher is an iconic mascot!")

	file.Write(text)

	// shorter way
	text = append(text, []byte(" SHORTER WAY")...)

	err = ioutil.WriteFile(fileName, text, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
