package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Worker struct {
	Id int
}

func NewWorker(id int) Worker {
	return Worker{Id: id}
}
func (w *Worker) ReadData(ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	val, ok := <-ch
	select {
	case <-ctx.Done():
		fmt.Println("CLOSED ")
	default:
		if ok {
			fmt.Printf("Worker %v recived data : %v\n", w.Id, val)
		} else {
			fmt.Print("Chan CLOSED")
			return
		}

	}
	defer wg.Done()
}

func main() {
	ch := make(chan int)
	pool := 5
	workers := make([]Worker, pool)
	var wg sync.WaitGroup
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-ctx.Done():
			close(ch)
			os.Exit(1)
		default:
			go func() {
				defer wg.Done()
				for {
					ch <- rand.Intn(100)
				}
			}()

			for i := 0; i < pool; i++ {
				workers[i] = NewWorker(i)
				wg.Add(1)
				go workers[i].ReadData(ch, &wg, ctx)
			}
			wg.Wait()
		}
	}

}
