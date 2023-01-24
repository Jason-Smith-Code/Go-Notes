package main

import "fmt"

// Maps

func main() {
	// creating a map
	// variable := map[type of Key]Type of value{enter values here}
	map1 := map[string]int{
		"john": 42,
		"brian": 12,
		"Sue" : 52,
	}
	// print map
	fmt.Println(map1)
	// print a specific value
	fmt.Println(map1["Sue"]) // prints 52

	// re assign value in map
	map1["brian"] = 80

	// deleting an element from the map : delete(mapName, key)
	delete(map1, "john")

}