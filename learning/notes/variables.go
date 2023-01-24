// Variables

package main

import (
	"fmt"
)

// Variable declaration : use the Var keyword to declare variables OR use := (this way you have to assign a value)

// var : Can be used inside and outside of functions, Variable declaration and value assignment can be done separately
// := : Can only be used inside functions, Variable declaration and value assignment cannot be done separately (must be done in the same line)
// const : if the variable is not expected to change, accessable from whole package

// Variable Types : You always have to specify either type or value (or both).
var integer int // whole number, positive or negative, default value of 0
var floatNumber float32 // decimal number, positive or negative, default value of 0
var word string // a word, default value of ""
var trueFalse bool // true or false, default value of false

// Const


// multiline
const (
	first = 1
	second = 2
)

// const using iota
// resets to 0 in different constant blocks
// increments by one each line in a constant block
const (
	one = iota // defaults to 0
	two = iota // defaults to 1 and so on
)

const (
	one1 = iota + 6 // 6
	two2 = iota // inherits previous const +6, and base is now 1, so equals 7
)


func variables() {
	integer2 := 2 // whole number, positive or negative
	floatNumber2 := 2.55 // decimal number, positive or negative
	word2 := "a word" // a word	
	trueFalse2 := true // true or false	

	fmt.Println(integer2)
	fmt.Println(floatNumber2)
	fmt.Println(word2)
	fmt.Println(trueFalse2)

	//  Pointers : stores the location in memory rather than the data itself
	firstName := "jason"
	fmt.Println(firstName) // prints variable

	ptr := &firstName
	fmt.Println(ptr) // prints point in memory
	fmt.Println(*ptr) // prints data in memory location

	firstName = "Sam"
	fmt.Println(ptr) // prints point in memory
	fmt.Println(*ptr) // prints the new data in same memory location
}


func main() {
	variables() // call the function
	fmt.Println(two2)
}
