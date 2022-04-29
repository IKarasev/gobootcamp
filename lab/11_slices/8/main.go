package main

import "fmt"

func main() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(newYears, years[:3]...)
	newYears = append(newYears, years[len(years)-3:]...)
	fmt.Println(newYears)
}
