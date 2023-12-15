package main

import (
	"log"
)

func main() {

	var mat [][]int = [][]int{
		{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
	}
	result := numSpecial(mat)
	log.Printf("result: %d\n", result)
	log.Printf("mat: %d\n", mat)
}

func numSpecial(mat [][]int) int {
	counter := 1
	var idx []int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			//вот мы проходим каждый i,j
			if mat[i][j] == 1 {
				idx = append(idx, j)
			}
		}
	}
	unique := map[int]bool{}
	// idx = all i where contains 1
	// for all idx if we have dublicate - ignore^ else
	return counter
}

/*

0,0 = 1
1,0 = 0
2,0 = 0

когда по одному i для всех j только одно совпадание 1
*/
