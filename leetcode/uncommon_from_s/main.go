package main

// TODO : https://leetcode.com/problems/uncommon-words-from-two-sentences/description/

import "log"

func main() {

	result := uncommonFromSentences("this apple is sweet", "this apple is sour")
	log.Println(result)

}

func uncommonFromSentences(s1 string, s2 string) map[string]int {
	var counter map[string]int
	counter = make(map[string]int)
	splitSentence(s1, counter)
	splitSentence(s2, counter)
	return counter
}

func splitSentence(s1 string, counter map[string]int) {
	wordSt := 0
	wordEn := 0
	for i := 0; i < len(s1); i++ {
		//раделить слова на вектора. Признак слова пробелы по бокам
		if s1[i] == ' ' {
			wordEn = i
			if _, ok := counter[s1[wordSt:wordEn]]; !ok {
				counter[s1[wordSt:wordEn]] = 0
			} else {
				counter[s1[wordSt:wordEn]]++
			}
			wordSt = i + 1
		}
	}
	if _, ok := counter[s1[wordSt:wordEn-1]]; !ok {
		counter[s1[wordSt:wordEn-1]] = 0
	} else {
		counter[s1[wordSt:wordEn-1]]++
	}

}

/*
A sentence is a string of single-space separated words where each word consists only of lowercase letters.

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.
*/
