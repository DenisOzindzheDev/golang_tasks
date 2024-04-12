package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, len(arr))

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ch <- val * val

		}(arr[i])
	}
	wg.Wait()
	close(ch)
	result := 0

	for v := range ch {
		result += v
	}

	fmt.Print(result)
}
