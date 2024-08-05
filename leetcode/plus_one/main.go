package main

func main() {

}
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0

	}
	digits = append([]int{1}, digits...)
	return digits
}

/*
[1,2,3]
1 9 9 + 1
2 0 0

1 2 3
-----------
100 1 * 100   i=0
20  2 * 10	  i=1
3   3 * 1     i=2

len 3 = 10 * 10 (len - 1 times)
len 2 = 10
len 1 = 1
-----------

*/
