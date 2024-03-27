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

/*
func findDuplicates(nums []int) []int {
    output := []int{}
    for _, num := range nums {
        idx := abs(num)
        if nums[idx-1] < 0 {
            output = append(output, idx)
        } else {
            nums[idx-1] *= -1
        }
    }
    return output
}
*/
