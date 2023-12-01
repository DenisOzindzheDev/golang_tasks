package main

import "log"

func main() {
	//var nums []int
	//solution := twoSum(nums, 5)
	//idxMap := make(map[int]int)
	//fmt.Println(solution)
	test := twoSum([]int{1, 2, 3, 4, 5, 6}, 3)
	log.Printf("testing array %d", test)

}

func twoSum(nums []int, target int) []int {
	//solution := make([]int, 5)
	idxMap := make(map[int]int)
	if len(nums) >= 2 && len(nums) <= 104 {
		if target >= -109 && target <= 109 {
			for idx, num := range nums {
				if v, found := idxMap[target-num]; found {
					return []int{v, idx}
				}
				idxMap[num] = idx
			}

		}
	}
	return []int{idxMap[0], idxMap[1]}
}
