package main

import "fmt"

/*var name string
var location string
var age int*/

/*var (
    name string
    location string
    phone_no int
)*/

/*
var (
	name     = "Ashish"
	location = "toronto"
	phone_no = 78678989
)

var (
	emp_name, emp_location, emp_phone_no = "Ashish", "toronto", 56789
)
*/

const (
	PI    = 3.14
	Truth = false
	club  = "Barcelona"
	Big   = 1 << 62 // I do not know how this works
)

func main() {
	//fmt.Println("Hello, World!")
	name, location := "Ashish", "Toronto"
	age := 32
	fmt.Println()
	//fmt.Printf("%s age %d from %s", name, age, location)
	fmt.Printf("%s age %d from %s ", name, age, location)
	fmt.Println()

	action := func() {
		fmt.Println("I am an anonymous function")

	}
	action()
	fmt.Println(Big)
}
