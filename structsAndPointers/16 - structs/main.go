package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{
		firstName: "Alex", lastName: "Dopeness",
	}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
