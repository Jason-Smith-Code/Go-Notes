// Variables

package main

import (
	"fmt"
)

// Variable declaration : use the Var keyword to declare variables OR use := (this way you have to assign a value)

// var : Can be used inside and outside of functions, Variable declaration and value assignment can be done separately
// := : Can only be used inside functions, Variable declaration and value assignment cannot be done separately (must be done in the same line)
// const : if the variable is not expected to change

// Variable Types : You always have to specify either type or value (or both).
var integer int // whole number, positive or negative, default value of 0
var floatNumber float32 // decimal number, positive or negative, default value of 0
var word string // a word, default value of ""
var trueFalse bool // true or false, default value of false

func variables() {
	integer2 := 2 // whole number, positive or negative
	floatNumber2 := 2.55 // decimal number, positive or negative
	word2 := "a word" // a word	
	trueFalse2 := true // true or false	

	fmt.Println(integer2)
	fmt.Println(floatNumber2)
	fmt.Println(word2)
	fmt.Println(trueFalse2)
}


func main() {
	variables() // call the function
}