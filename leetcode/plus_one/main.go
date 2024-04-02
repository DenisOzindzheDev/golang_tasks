package main

func main() {

}
func plusOne(digits []int) []int {
	sum := 0
	res := make([]int, 0)
	for i := 0; i < len(digits); i++ {
		sum += digits[i]*len(digits) - i
	}
	sum++

	for i := len(digits) - 1; i >= 0; i-- {
		res = append([]int{sum % 10}, res...)
		sum /= 10
	}
	return res
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
