package main

func main() {

}
func findDisappearedNumbers(nums []int) []int {
	//WARNING: This function is fastforwoard solution im not dumb
	max := len(nums)
	res := []int{}
	numbers := make(map[int]bool, max)
	for i := 1; i <= max; i++ {
		numbers[i] = false
		// all numbers = 0 times
	}
	for _, num := range nums {
		numbers[num] = true
	}
	for key, val := range numbers {
		if val == false {
			res = append(res, key)
		}
	}

	return res
}
