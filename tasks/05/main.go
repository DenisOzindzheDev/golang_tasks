package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/net/context"
)

func main() {
	timeoutSec := time.Millisecond * 15000
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSec)
	var mutex = &sync.Mutex{}
	ch := make(chan int)

	defer cancel()

	for {
		select {
		case <-ctx.Done():
			close(ch)
		default:
			go func(ctx context.Context) {

				fmt.Println("Claimed message")
				select {
				case <-ch:
					mutex.Lock()
					ch <- rand.Intn(10)
					mutex.Unlock()
				default:
					fmt.Printf("Channel closed")
				}
				time.Sleep(time.Second * 1)
			}(ctx)
			go func(ctx context.Context) {
				if val, ok := <-ch; ok {
					fmt.Println("recived ", val)
				}
			}(ctx)
		}

	}
	//close(ch)

}
