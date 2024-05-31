package main

func main() {

}

func singleNumber(nums []int) int {
	visited := make(map[int]int, 0)
	for _, i := range nums {
		visited[i]++
	}
	for key, val := range visited {
		if val == 1 {
			return key
		}
	}
	return 0
}
