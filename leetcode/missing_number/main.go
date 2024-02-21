package main

import (
	"sort"
)

func main() {

}
func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			return nums[i] + 1
		}
	}
	return nums[len(nums)-1] + 1
}
