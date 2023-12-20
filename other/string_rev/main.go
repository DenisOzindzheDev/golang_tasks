package main

import (
	"fmt"
)

func main() {
	fmt.Printf("result %v", Reverse("Welcome to the reverse"))
}

func Reverse(s string) string {
	var r []rune
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, rune(s[i]))
	}
	return string(r)
}
