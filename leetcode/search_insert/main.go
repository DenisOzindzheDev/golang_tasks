package main

import "log"

func main() {
	test := searchInsert([]int{1, 3}, 2) // need 1
	log.Printf("Searching results for test %d", test)

}

// You must write an algorithm with O(log n) runtime complexity.
func searchInsert(nums []int, target int) int {
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == target {
			return i
		} else if target > nums[i] && target < nums[i+1] {
			return i + 1
		}
	}
	return 0
}

// find if we have match
// if match doesn't exist find i where we closer to element
// если мы не нашли эллемент значит он произвольно в любом месте близок к одном из значений, надо найти место где таргет > and <  рядом элементов и ввернуть i
// текущий эллемент не является таргетом и он меньше таргета и следующий эллемент больше таргета
