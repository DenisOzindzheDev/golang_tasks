package main

func findDuplicates(nums []int) []int {
	numbers := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}
	var duplicates []int
	for i := 0; i < len(nums); i++ {
		numbers[nums[i]]++
		if numbers[nums[i]] > 1 {
			duplicates = append(duplicates, nums[i])
		}
	}

	return duplicates
}
