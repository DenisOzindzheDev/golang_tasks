package main

import "log"

func main() {
	res := rearrangeArray([]int{3, 1, -2, -5, 2, -4})
	log.Printf("Result of rearrange %v", res)
}
func rearrangeArray(nums []int) []int {
	var pos, neg, res []int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			pos = append(pos, nums[i])
		} else {
			neg = append(neg, nums[i])
		}
	}
	l1 := 0
	l2 := 0
	for i := 0; i < len(nums); i++ {

		if i%2 == 0 {
			res = append(res, pos[l1])
			l1++
		} else {
			res = append(res, neg[l2])
			l2++
		}
	}
	return res
}
