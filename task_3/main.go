package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	var (
		arr  = [5]int{2, 4, 6, 8, 10}
		wg   sync.WaitGroup
		ch   chan int = make(chan int) //non buffered channel
		arr2 [5]int                    //int of array
	)

	for i := 0; i <= 4; i++ {
		wg.Add(1)
		i := i
		go func(val int, wg *sync.WaitGroup, ch chan int) {
			fmt.Printf("result: %d\n", val*val)
			ch <- val * val //int
			arr2[i] = <-ch
			defer wg.Done()

		}(arr[i], &wg, ch)

	}

	wg.Wait()
	defer close(ch)

	//wg.Done()
}

// 1 Wait Group
// 1 routrne make logic - done
// 2 routrne make logic = done
// 3 routine make logic - done
// Wait for all goroutines to finish
