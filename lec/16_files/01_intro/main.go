package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const lecHeader string = "Files: base operations"

func main() {
	pln := fmt.Println
	pf := fmt.Printf
	pf("%s\n", lecHeader)
	pln(strings.Repeat("-", 10))

	// creating file
	var newF *os.File // creting pointer to new file
	pf("%T:%v\n", newF, newF)

	// no exceptions in go!
	// can use error type
	var err error
	newF, err = os.Create("a.txt") // creates new file in cur dir, if file exists - truncates it

	// if error - exit
	if err != nil {
		// pln(err)   // print error msg
		// os.Exit(1) // exit program with code 1

		// another way:
		log.Fatal(err) // prints error and exits
	}

	pf("%T:%v\n", newF, newF)
	pf("error: %v\n", err)

	// truncate file
	err = os.Truncate("a.txt", 0) // 0 - complitely empty file, otherwise truncate to n bites
	if err != nil {
		log.Fatal(err)
	}

	// after done with file - close it !!!
	newF.Close()
	pf("%T:%v\n", newF, newF)

	// open file
	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err) // => The system cannot find the file specified, if file not exist
	}
	pf("%T:%v\n", file, file)
	file.Close()

	// open with options
	file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
	//os.OpenFile("a.txt", os.O_APPEND, 0644)
	// file name\path, open mode, permissions
	if err != nil {
		log.Fatal(err) // => The system cannot find the file specified, if file not exist
	}
	pf("%T:%v\n", file, file)
	file.Close()

	pln(strings.Repeat("-", 10))

	// getting file info
	var fInfo os.FileInfo
	fInfo, err = os.Stat("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	pf("%T\n", fInfo)
	pln(fInfo) // => &{a.txt 32 {426084120 30957478} {1728439882 30957479} {1728439882 30957479} 0 0 0 0 {0 0} E:\Cources\Udemi\Go\Go_Bootcamp\practice\lec\16_files\01_intro\a.txt
	//0 0 0 false}
	pln("name:", fInfo.Name())
	pln("size in bytes:", fInfo.Size())
	pln("last mod:", fInfo.ModTime())
	pln("is dir:", fInfo.IsDir())
	pln("permissions:", fInfo.Mode())

	pln(strings.Repeat("-", 10))

	// check if file exists
	fInfo, err = os.Stat("ab.txt")
	pln("file \"ab.txt\":", err)
	if err != nil {
		if os.IsNotExist(err) {
			pln("file \"ab.txt\" does not exist")
		}
	}
	_ = fInfo

	pln(strings.Repeat("-", 10))

	// move\rename
	oldPath := "a.txt"
	newPath := "b.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	// remove file
	err = os.Remove("b.txt")
	if err != nil {
		log.Fatal(err)
	}

}
