package main

import "log"

func main() {
	test := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	solution := removeDuplicates(test)
	log.Printf("solution: %d", solution)
}
func removeDuplicates(nums []int) int {
	unique := 1
	matched := nums[0]
	for i := 1; i < len(nums); i++ {
		if matched != nums[i] {
			unique++
			matched = nums[i]
		}
	}
	return unique
}
