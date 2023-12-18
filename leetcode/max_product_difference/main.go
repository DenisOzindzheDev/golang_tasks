package main

import (
	"log"
	"sort"
)

func main() {

	test := maxProductDifference([]int{5, 6, 2, 7, 4})
	log.Printf("Result of product %d ", test)

}
func maxProductDifference(nums []int) int {

	sort.Ints(nums)

	return (nums[len(nums)-1] * nums[len(nums)-2]) - (nums[0] * nums[1])
}
