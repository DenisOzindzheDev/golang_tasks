package main

import (
	"log"
)

func main() {
	res := romanToInt("XIV")
	log.Printf("%v", res)

}

var (
	values = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
	}
)

func romanToInt(s string) int {
	result := 0

	for i := len(s) - 1; i > 0; i-- {
		if values[string(s[i-1])] < values[string(s[i])] {
			result -= values[string(s[i-1])]
		} else {
			result += values[string(s[i-1])]
		}
	}
	result += values[string(s[len(s)-1])]

	return result
}
