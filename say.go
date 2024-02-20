// Package say returns friendly statements
package say

import v2 "github.com/cadamus/say/v2"

// Hello returns the famous "hello, World!" greeting
func Hello() string {
	return v2.Hello("")
}

// Goodbye returns the infomous Lawrence Oates quote
func Goodbye() string {
	return v2.Goodbye()
}
