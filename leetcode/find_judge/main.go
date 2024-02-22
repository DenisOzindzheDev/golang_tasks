package main

import (
	"log"
)

func main() {
	res := findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}})
	log.Printf("found %v", res)

}

func findJudge(n int, trust [][]int) int {
	peoples := make(map[int]bool)
	trustFactor := make(map[int]int)

	for i := 0; i < len(trust); i++ {
		peoples[trust[i][0]] = true
		trustFactor[trust[i][1]] += 1
	}
	for i := 1; i <= n; i++ {

		if trustFactor[i] == n-1 && !peoples[i] {
			return i
		}
	}
	return -1
}

//array, hash table, graph

/*
	Прищнаки судьи :
	1) У него самый высокий траст фактор <-- legacy
	2) Доверие к нему len()-1
	3) У него 0ой счетчик доверия к другим

	1 3
	2 3
	3 1

	3 ++
	3 ++
	1 ++

	3 match = len -1
	3 pairs 1

	false

	Чтобы обозначить город: (hash)
	1) Пройтись по всем парам и установить для пары true если у нее есть пара
	2) Для пары у которой нету match посчитать count попаданий

*/

// [[1,3],[1,4],[2,3],[2,4],[4,3]]
