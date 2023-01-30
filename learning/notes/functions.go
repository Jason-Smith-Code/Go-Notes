package main

import "fmt"

// creating a function : func functionName(arguments) {}
// a function in GO, can return multiple values

// func main() : the entry point of the application

func main() {
	// invoke a deferred function : this will run last, if there is more than 1 deferred function, they will run in the order they have been deferred.
	defer functionName()

	// invoke functions
	port := 3000
	port, err := startWebServer(port)
	fmt.Println(err)

}

// func nameOfFunction(argument argumentType) (returnDataType, returnDataType, returnDataType)
func startWebServer(port int) (int, error) { // int + error refers to the data type returned by the function
	// return data
	return port, nil
}

func functionName() {

}


// functions with receivers : func (receiverVariable receiver) functionName() {}
