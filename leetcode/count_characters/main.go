package main

import "strings"

func main() {

}
func countCharacters(words []string, chars string) int {
	char := map[rune]bool{}
	result := 0
	for _, c := range chars {
		char[c] = true
	}
	for i := 0; i < len(words)-1; i++ {
		for _, c := range words[i] {
			if char[c] != strings.Contains(string(chars[i]), string(c)) {
				break
			}
			result += len(words[i])
		}
	}
	return result
}
