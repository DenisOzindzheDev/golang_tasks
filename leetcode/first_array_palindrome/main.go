package main

import "log"

func firstPalindrome(words []string) string {
	for _, word := range words {
		if check(word) {
			return word
		}
	}
	return ""
}
func check(word string) bool {
	log.Println("Checking if words %v and %v are same", word, reverse(word))
	if word == reverse(word) {
		log.Print("SUM")
		return true
	}
	return false
}
func reverse(word string) string {
	var res []byte
	for i := 0; i < len(word); i++ {
		res = append(res, byte(word[len(word)-i-1]))
	}
	log.Printf("reverse called with %v", string(res))
	return string(res)
}
func main() {
	res := firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"})
	log.Printf(res)
}
