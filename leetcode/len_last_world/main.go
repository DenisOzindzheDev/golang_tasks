package main

func main() {

}
func lengthOfLastWord(s string) int {
	// Оборвать конец фразы пробел может быть произваольной длины
	first := len(s) - 1 // index of last space
	for s[first] != ' ' {
		first--
	}
	last := first // index of last space

	for last > 0 && s[last-1] != ' ' {
		last--
	}
	return first - last - 1
}

// c := 0

// for i := len(s) - 1; i >= 0; i-- {

// 	if s[i] == ' ' {
// 		continue
// 	}
// 	c++
// 	if s[i] != ' ' && s[i-1] == ' ' {
// 		break
// 	}

// }

// return c
