package main

import "fmt"

// Slice
// can be resized


func main() {
	// create an slice the short way
	slice := []int{1, 2, 3} // same as array but no size declaration
	fmt.Println(slice) // slice is [1, 2, 3]
	// add item or items to slice
	slice = append(slice, 4)
	fmt.Println(slice) // slice is now [1, 2, 3, 4]
	slice2 := slice[1:] // creates a new slice from another slice using part of the other slice[start:end]
	fmt.Println(slice2) // slice is now [2, 3, 4] index started 1, so we miss out 0 index and print the rest
	slice3 := slice[:1] // creates a new slice from another slice using part of the other slice[start:end(not including)]
	fmt.Println(slice3) // slice is now [1] because it only goes from index 0, then ends with and not including index 1
}