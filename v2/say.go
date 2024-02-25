// Package say returns friendly statements
package say

import "fmt"

// Hello returns the "hello, {name}!" greeting.  If name is the empty string, name will
// default to World, returning the famous hello, World! greeting.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("hello, %v!", name)
}

// Goodbye returns the infomous Lawrence Oates quote
func Goodbye() string {
	// return "I am just going outside and I may be some time"
	return "This is not a goodbye, my darling, this is a thank you."
}
