package main

func main() {
}
func containsDuplicate(nums []int) bool {
	c := make(map[int]int, 0)
	for _, i := range nums {
		c[i]++
		if c[i] > 1 {
			return true
		}
	}

	return false

}
