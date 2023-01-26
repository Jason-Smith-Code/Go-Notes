package main

import "fmt"

// creating a function : func functionName(arguments) {}

// func main() : the entry point of the application

func main() {
	// invoke functions
	port := 3000
	port, err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) { // int + error refers to the data type returned by the function
	// return data
	return port, nil
}