package main

import (
	"fmt"
	"strings"
)

func main() {
	i := strStr("leetcode", "leet")
	fmt.Println(i)
}

func strStr(haystack string, needle string) int {
	//var characters []rune = []rune(haystack)
	for i := 0; i < len(haystack); i++ {
		if ok := strings.HasPrefix(haystack[i:len(haystack)], needle); ok {
			return i
		}
	}
	return -1
}

//a
