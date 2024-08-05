package main

func main() {

}

func mySqrt(x int) int {
	left, right := 0, x+1

	for left < right {
		mid := left + (right-left)/2

		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left - 1

}

// func binarySearch(arr []int, f, l, h int) int {
// 	for l <= h {
// 		m := l + (h-l)/2
// 		if arr[m]*arr[m] == f {
// 			return arr[m]
// 		}
// 		if arr[m]*arr[m] > f {
// 			l = m + 1
// 		} else {
// 			h = m - 1
// 		}

// 	}
// 	return 0
// }
