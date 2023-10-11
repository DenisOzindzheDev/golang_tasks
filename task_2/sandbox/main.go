package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Printf("Done")
	var wg sync.WaitGroup
	wg.Add(2)

	go mock(1, &wg)
	go mock(2, &wg)

	wg.Wait()

}

func mock(id int, wg *sync.WaitGroup) {
	fmt.Printf("routine %d started \n", id)
	time.Sleep(time.Second * 5)
	fmt.Printf("routine %d ended \n", id)
	wg.Done()
}
