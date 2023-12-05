package main

import "log"

func main() {
	sol := numberOfMatches(1)
	log.Printf("res %v", sol)
}

func numberOfMatches(n int) int {
	matches := 0
	for n > 1 { //while teams more than 1 (winner)
		if n%2 == 0 {
			matches += n / 2
			n /= 2
		} else {
			matches += n/2 + 1
			n /= 2
		}
	}
	return matches
}

/*
7 teams n
round 1 (6 teams) (teams - 1) - condition teams !odd
round 2 (3 teams) (teams - half of round 1)+ 1 - conditon odd

*/
