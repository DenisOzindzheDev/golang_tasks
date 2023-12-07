package main

func main() {

}

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i]%2 == 0 {
			num = num[:i]
		} else {
			break
		}
	}
	return num
}
