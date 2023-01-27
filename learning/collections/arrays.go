package main

import "fmt"

// Array
// all elements must have same type
// array size must be declared and cannot be resized


func main() {
	// create an array the long way
	var array1 [3]int

	// assign data to an array
	array1[0] = 1
	array1[1] = 1
	array1[2] = 1

	fmt.Println((array1)) // print array
	fmt.Println((array1[1])) // print an item from an array

	// create an array the short way
	array2 := [3]int{1, 2, 3}
	fmt.Println((array2)) // print array
	fmt.Println((array2[1])) // print an item from an array

}