package main

func firstMissingPositive(nums []int) int {
	len := len(nums)
	for i := 0; i < len; i++ {
		if nums[i] <= 0 || nums[i] > len {
			nums[i] = len + 1
		}
	}

	for i := 0; i < len; i++ {
		val := abs(nums[i])
		if val >= 1 && val <= len {
			idx := val - 1
			if nums[idx] > 0 {
				nums[idx] *= -1
			}
		}
	}
	for i := 1; i <= len; i++ {
		if nums[i-1] > 0 {
			return i
		}
	}
	return len + 1

}

// absolute number
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// for _, num := range nums {
// 	idx := abs(num)
// 	if nums[idx-1] < 0 {
// 		output = append(output, idx)
// 	} else {
// 		nums[idx-1] *= -1
// 	}
/* SOLUTION MINDMAP */
// i        |
// a[3,4,-1,1]

//  smallest = 3
//  next = 4

//  smallest = 1
//  next = 2

//  next missing return next
