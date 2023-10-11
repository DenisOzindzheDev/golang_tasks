package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go square(arr[i], &wg)
		wg.Done()
	}
	wg.Wait()

}

func square(val int, wg *sync.WaitGroup) {
	fmt.Println(val * val)
	wg.Done()
}
