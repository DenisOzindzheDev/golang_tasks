package main

func main() {

}
func findMaxK(nums []int) int {
	m := make(map[int]int, len(nums)) //buffered map
	for k := range m {
		for i := range nums {
			m[k] = i
			if -k == i {
				return k
			}
		}
	}
	return 0
}

/*
Example 1:

Input: nums = [-1,2,-3,3]
Output: 3
Explanation: 3 is the only valid k we can find in the array.
*/
