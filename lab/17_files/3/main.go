package main

import (
	"log"
	"os"
)

func main() {
	fileName := "data.txt"

	err := os.Remove(fileName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("File \"%s\" removed\n", fileName)
	}
}
