package main

import "log"

func main() {

	test := imageSmoother([][]int{{}, {}, {}})
	log.Printf("TEST image %v", test)
}
func imageSmoother(img [][]int) [][]int {
	resultMatrix := make([][]int, len(img))
	for v := range img {
		resultMatrix[v] = make([]int, len(img[0]))
	} //Создаем результирующую матрицу размером img

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ { // Проход по матрице

			resultMatrix[i][j] = 1 //Что-то происходит
		}

	}

	return resultMatrix
}


/*
0 0 0   для 1 выборка  с {i j}
0 1 0    
0 0 0
*/