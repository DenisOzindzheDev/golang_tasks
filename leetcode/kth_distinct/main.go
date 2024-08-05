package main

func main() {

}

func kthDistinct(arr []string, k int) string {
	values := make(map[string]int, len(arr)) //all
	for _, v := range arr {
		values[v]++
	}
	c := 0
	for _, key := range arr {
		if values[key] == 1 {
			if c == k-1 {
				return key
			} else {
				c++
			}
		}
	}
	return ""
}
