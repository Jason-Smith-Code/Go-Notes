package main

import "fmt"

// Structs
// can have a mixture of types of data stored in a struct

func main() {
	// creating a struct
	// define struct
	type user struct {
		// enter fields 
		ID int // defaults to 0
		FirstName string // defaults to ""
		LastName string // defaults to ""
	}
	// initialise struct
	var person user

	// add data to a struct field 
	person.ID = 1
	person.FirstName = "jason"
	person.LastName = "Smith"
	// or
	person2 := user{ ID: 2, FirstName: "betty", LastName: "boo"}



	fmt.Println(person)
	fmt.Println(person2)
}