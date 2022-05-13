package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {
	var myCar vehicle = car{
		licenceNo: "aa123bb",
		brand:     "Audi",
	}

	fmt.Println("Car name:", myCar.Name())
	fmt.Println("licence №", myCar.License())
}
