package main

func main() {

}
func removeElement(nums []int, val int) int {
	for i := len(nums) - 1; i > -1; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
