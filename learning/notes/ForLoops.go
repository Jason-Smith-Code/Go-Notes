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

	// looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
	oldMap := map[string]int{
		"john": 42,
		"brian": 12,
		"Sue" : 52,
	}
	newMap := map[string]int {}

	for key, value := range oldMap {
		newMap[key] = value
	}
}