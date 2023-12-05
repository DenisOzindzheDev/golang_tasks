package main

import (
	"log"
)

func main() {
	result := isValid("()") // true
	log.Printf("res %v", result)
}

var mappings = map[byte]byte{
	'{': '}',
	'(': ')',
	'[': '}',
}

func isValid(s string) bool {
	var charstack []byte

	for _, c := range s { // i = 0
		if _, ok := mappings[byte(c)]; ok {
			charstack = append(charstack, byte(c)) // [  { ]
		} else if len(charstack) == 0 || mappings[charstack[len(charstack)-1]] != byte(c) {
			return false
		} else {
			charstack = charstack[:len(charstack)-1]
		}
	}

	return len(charstack) == 0
}

/*
Hint 1
Use a stack of characters.
Hint 2
When you encounter an opening bracket, push it to the top of the stack.
Hint 3
When you encounter a closing bracket, check if the top of the stack was the opening for it. If yes, pop it from the stack. Otherwise, return false.
*/
