// strings Package

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Clone : func Clone(s string) string

	// Compare : func Compare(a, b string) int
	fmt.Println(strings.Compare("a", "e")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("b", "a")) // 1

	// Contains : func Contains(s, substr string) bool
	// contains(string, substring) returns true or false
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", "")) // true
	fmt.Println(strings.Contains("", "")) // true
}