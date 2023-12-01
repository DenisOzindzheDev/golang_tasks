package main

import (
	"fmt"
	"log"
)

func main() {

	test := isPalindrome(2312)
	fmt.Println(test)
}

func isPalindrome(x int) bool {
	if x >= 0 {
		return x == reverse_int(x)
	}
	return false
}

func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		log.Printf("remainder %d", remainder)
		new_int *= 10
		log.Printf("new_int %d", new_int)
		new_int += remainder
		log.Printf("new_int %d", new_int)
		n /= 10
		log.Printf("n %d", n)
	}
	return new_int
}
