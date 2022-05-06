package main

import "fmt"

type Grades struct {
	grade  int
	course string
}

type Person struct {
	name   string
	age    int
	colors []string
	grades Grades
}

func main() {
	me := Person{
		name:   "My name",
		age:    30,
		colors: []string{"black", "white"},
		grades: Grades{
			grade:  95,
			course: "Lang",
		},
	}
	you := Person{
		name:   "Your name",
		age:    32,
		colors: []string{"green", "blue"},
		grades: Grades{
			grade:  91,
			course: "Math",
		},
	}

	me.grades.course = "Golang"
	me.grades.grade = 98

	fmt.Printf("me = %#v\n", me)
	fmt.Printf("%+v\n", me)
	fmt.Printf("you = %#v\n", you)
	fmt.Printf("%+v\n", you)
}
