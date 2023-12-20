package main

import "sort"

func buyChoco(prices []int, money int) int {
	matches := 0
	totalSpent := money // 4

	sort.Ints(prices) // 1 1
	for _, p := range prices {
		if matches <= 2 { // 0
			if p >= totalSpent && matches != 1 { // 1 !> 4
				return money
			} else { // then
				totalSpent -= p // 4 -1  = 3
				matches++       // 1 match!
			}
		} else {
			break
		}

	}
	return totalSpent
}

func buyChoco2(prices []int, money int) int {
	sort.Ints(prices)
	if r := prices[0] + prices[1]; r <= money {
		return money - r
	}
	return money
}
