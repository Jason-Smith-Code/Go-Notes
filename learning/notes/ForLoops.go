package main

// For loops : there are 3 types of for loop
// for init; condition; post { }
// for condition { }
// infinite loop : for { }

func main() {
	// for loop example
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// looping over a map, a range clause can manage the loop.
	for key, value := range oldMap {
		newMap[key] = value
	}

	
	for key := range oldMap {
		if key.expired() {
			delete(oldMap, key)
		}
	}

	sum := 0
	for _, value := range array {
		sum += value
	}
}