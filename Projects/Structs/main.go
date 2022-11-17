package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//var Alex person
	//Alex := person{firstName: "Alex", lastName: "Anderson"}
	jim := person{
		firstName: "Jim",
		lastName:  "Carry",
		contact: contactInfo{
			email:   "jimcarry@hollywood.com",
			zipcode: 123445,
		},
	}
	jim.updateFirstName("James")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
