package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const lecHeader string = "Files: buffered writer!"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// using bufio lets us write to file after buffering text in memory
	// if a lot of manipulations with data is needed
	// or if we need to write byte by byte

	file, err := os.OpenFile(
		"a.txt",
		os.O_WRONLY|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// init writer
	bufWriter := bufio.NewWriter(file)

	bs := []byte{90, 91, 92, 93, 94, 95}

	bWritten, err := bufWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}

	pln("Bytes written to BUFFER!:", bWritten)

	bAvailable := bufWriter.Available()            // gives available bytes to use
	pln("Bytes availablle in buffer:", bAvailable) // => 4090, 4096 - default budder size, 6 used

	// we can write str directly to buffer
	bWritten, err = bufWriter.WriteString("\nMany other words\nand more")

	if err != nil {
		log.Fatal(err)
	}

	pln("Bytes of string written to BUFFER!:", bWritten)

	bAvailable = bufWriter.Available()
	pln("Bytes availablle in buffer:", bAvailable) // => 4064, 4096 - default budder size, used: 6 before + new 26 bytes

	bUnflushed := bufWriter.Buffered()  // => # of bytes in buffer
	pln("Bytes in byffer:", bUnflushed) // => 32

	// write buffer to file
	bufWriter.Flush()

	// reset all non flushed buffer changes
	bufWriter.Reset(bufWriter)

}
