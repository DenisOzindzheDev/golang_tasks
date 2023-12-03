package main

import (
	"log"
	"strings"
)

func main() {
	testData := []string{"", "flow", "flight"}
	result := longestCommonPrefix(testData)
	log.Printf("%s\n", result)

}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs {
		for !strings.HasPrefix(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
