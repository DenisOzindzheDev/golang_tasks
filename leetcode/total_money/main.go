package main

import "log"

func main() {
	test := totalMoney(4) //amount of money
	log.Printf("res = %v", test)
}

func totalMoney(n int) int {
	currMoney := 1
	weekIdx := 0
	for i := 0; i < n; i++ {
		currMoney += i + weekIdx
		if i%7 != 0 {
			weekIdx++
		}
	}

	return currMoney
}

/*
Simulate the process
by keeping track of how much money
John is putting in and which day of the week it is,
and use this information to deduce how much money
John will put in the next day.

N = n+n+1
1
1 + 1 + 1
3 + 3 + 1

n += i

i n
1 1
2 3
3 6
4 10
5 15
6 21
7 28

8 30
9 33

2  = - 5
3

*/
