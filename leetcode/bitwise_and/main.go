package main

import (
	"fmt"
)

func main() {
	input := 50
	fmt.Printf("input: %b\n", input)

}
func rangeBitwiseAnd(left int, right int) int {
	ans := left & right
	distance := right - left + 1

	for i := 0; i < 32; i++ {
		// Check if there's any zero bits within the range
		// at the corresponding position
		if (ans>>i)&1 == 1 && distance > (1<<i) {
			ans = ans ^ (1 << i)
		}
	}

	return ans
}
