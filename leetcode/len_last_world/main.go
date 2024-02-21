package main

import "strings"

func main() {

}
func lengthOfLastWord(s string) int {
	// Оборвать конец фразы пробел может быть произваольной длины
	substr := s
	for strings.HasPrefix(substr, " ") {
		substr, _ = strings.CutPrefix(substr, " ")

	}
	// Собрать слово до крайнего пробела и вернуть его длину

}
