package main

import (
	"fmt"
	"strings"
)

func main() {
	input, input2 := []string{"abc", "c"}, []string{"abcc"}
	eq := arrayStringsAreEqual(input, input2)
	fmt.Println(input, input2, eq)
}
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	words, words2 := strings.Join(word1, ""), strings.Join(word2, "")
	return words == words2
}
