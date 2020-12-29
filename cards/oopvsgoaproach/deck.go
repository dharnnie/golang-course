package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

// I did not realise I could create type(s) of primitive golang types (as in -> type deck []string)like I'd do with a struct (as in -> type User struct{}), peharps struct is also a golang primitive type

// Notes (Good practice)
// Always refer to the recievers with 1 or 2 characters that match the type
