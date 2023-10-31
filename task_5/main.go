package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	ch := make(chan int)
	keepAliveTimeout := time.Second * 20

	go func() {
		for {
			ch <- rand.Int()
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			fmt.Println("value read ", <-ch)
		}
	}()

	time.Sleep(keepAliveTimeout)
}
