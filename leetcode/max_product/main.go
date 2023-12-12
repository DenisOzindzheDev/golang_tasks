package main

func main() {

}

func maxProduct(nums []int) int {
	a, b := 1, 0

	for _, i := range nums {
		if i > a && a < b {
			a = i
		} else if i > b && b <= a {
			b = i
		}
	}

	return (a - 1) * (b - 1)

}
