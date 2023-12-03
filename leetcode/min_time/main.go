package main

import (
	"log"
	"math"
)

func main() {
	testData := [][]int{
		{1, 1},
		{3, 4},
		{-1, 0},
	}
	result := minTimeToVisitAllPoints(testData)
	log.Printf("result: %v", result)
	//log.Printf("test data %v", testData[1])

}

func minTimeToVisitAllPoints(points [][]int) int {

	t := 0
	pos := points[0]
	// points[0] = {1,1}
	// points[1] = {3,4}
	// points[2] = {-1, 0}
	for _, p := range points {
		t += int(math.Max(math.Abs(float64(p[0])-float64(pos[0])), math.Abs(float64(p[1])-float64(pos[1]))))
		pos = p
	}

	return t

}

/*
func extractValues(points []int)  []int{
	for i := range points {


		for pos[0] != points[i+1][0] && pos[1] != points[i+1][1] { // whilt this pos x and y do not equal x y from next step do ++
			for x := 0; x < 2; x++ {
				if points[i+1][x] != pos[x] {
					pos[x] += 1
				}
			}
			log.Printf("step %v pos", pos)
			t++
		}



	}
}
*/
