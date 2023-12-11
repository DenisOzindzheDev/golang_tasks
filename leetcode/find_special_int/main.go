package main

import (
	"log"
)

func main() {
	result := findSpecialInteger([]int{1, 2, 3})
	log.Printf("found %d special integers ", result)
}

func findSpecialInteger(arr []int) int {
	frequency := make(map[int]int)
	result := arr[0]
	percent := int(float64(len(arr)) * 0.25)

	for _, i := range arr {
		frequency[i]++
	}

	for key, count := range frequency {
		if count > percent {
			result = key
		}
	} //хотел заменить на бинарный поиск - но прочитал что для больших датасетов можно искать по мапе и это будет проще и быстрее o(log n) https://stackoverflow.com/questions/62201467/can-we-use-maps-for-searching-instead-of-binary-search

	return result
}

// what is 25 percent? is val * 0.25
