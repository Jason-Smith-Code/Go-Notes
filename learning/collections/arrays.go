package main

import "fmt"

// Array
// all elements must have same type
// array size must be declared and cannot be resized

func main() {
	// create an array the long way : var arrayName [size of array]type of data stored in array
	var arrayName [3]int

	// assign data to an array : arrayName[index] = value
	arrayName[0] = 1
	arrayName[1] = 1
	arrayName[2] = 1

	fmt.Println((arrayName)) // print array
	fmt.Println((arrayName[1])) // print an item from an array

	// create an array the short way : variableName := [size of array]{content of array}
	array2 := [3]int{1, 2, 3}
	fmt.Println((array2)) // print array
	fmt.Println((array2[1])) // print an item from an array

}