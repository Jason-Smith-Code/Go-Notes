package main

import (
	"errors"
	"fmt"
)

// creating a function : func functionName(arguments) {}
// a function in GO, can return multiple values

// func main() : the entry point of the application

func main() {
	// invoke a deferred function : this will run last, if there is more than 1 deferred function, they will run in the order they have been deferred.
	defer deferredfunctionName()

	// invoke functions
	port := 3000
	port, err := startWebServer(port)
	fmt.Println(err)

	// Anonymous functions : 
	func() string{
		name := "jason"
		return name
	}()

	// assigning a function to a variable
	anomymousFunction := func() string{
		name := "jason"
		return name
	}
	// initiate anomymousFunction
	anomymousFunction()

	// store return value of function : function must return a value
	returnValue := anomymousFunction()
	println(returnValue)
}

// returning multiple values from a function
// func nameOfFunction(argument argumentType) (returnDataType, returnDataType, returnDataType)
func startWebServer(port int) (int, error) { // int + error refers to the data type returned by the function
	// return data ) for both parameters
	return port, nil
}

func deferredfunctionName() {

}



// variatic function : use a spread operator on a parameter to create a slice of
func sum(values ...int) int { // can only use spread operator on last parameter
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// functions with receivers : func (receiverVariable receiver) functionName() {}


// _ : means ignore parameter


// PUBLIC functions : capitalize the function name
func PublicFunctionName() {}
// import the file path into the current file
// then use dot notation to access the function from another file : fileName.PublicFunctionName()


// PRIVATE functions : lowercase the function name
func privarteFunctionName() {} // this cannot be access from another file


// named return function : seting the function up so that we simply "return" without explicitly declaring what we are returning on the final line of code.
func namedReturn(parameter1, parameter2 int) (answer int, err error) { // we give the return values names in the last parenthisis accompanied by their types
	// assign the answer return value
	answer = parameter1 + parameter2
	// assign the error return value
	if answer > 10 {
		err = errors.New("number too high")
	}
	// only use return on final line
	return
}


// function with a method receiver
type StructName struct {
	beginner, intermediate, advanced string
}
// short version
func (mr StructName) String()	 string {
	return mr.advanced + mr.beginner + mr.intermediate
}

// functions as parameters

// returning functions from functions

// holding state in a function

// * 

// &