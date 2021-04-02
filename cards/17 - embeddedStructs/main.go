package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	daniel := person{
		firstName: "Daniel",
		lastName:  "Osineye",
		contactInfo: contactInfo{
			email:   "daniel.osineye@gmail.com",
			zipCode: "12345",
		},
	}
	daniel.updateFirstname("Danny")
	daniel.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstname(newFirstName string) {
	p.firstName = newFirstName
}
