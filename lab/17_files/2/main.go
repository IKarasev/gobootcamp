package main

import (
	"log"
	"os"
)

func main() {
	fileName := "info.txt"

	_, err := os.Stat(fileName)

	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename(fileName, "data.txt")

	if err != nil {
		log.Fatal(err)
	}
}
