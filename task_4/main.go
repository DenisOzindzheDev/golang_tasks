package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
type Workers struct {
	Amount chan int
}

func (w *Workers) InitWorker(id int) {
	go func() {
		for woker := range w.Amount {
			time.Sleep(time.Second * 1)
			res := fmt.Sprintf("Worker %d: %d", id, woker)
			fmt.Println(res)
		}
	}()
}
func Gen(ch chan int) {
	ch <- rand.Int()
}

// У нас есть базлвые структуры чтобы работать с программой, теперь нужно реализовать
func main() {
	var workersAmount int
	var shutdownTimeoutSec int = 1
	//var wg sync.WaitGroup
	workers := Workers{
		Amount: make(chan int),
	}
	fmt.Print("Enter workers amount : ")
	fmt.Scanln(&workersAmount)
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGINT, syscall.SIGTERM)
	for i := 0; i < workersAmount; i++ {
		workers.InitWorker(i + 1)
	}

	//loop:
	for {
		select {
		case <-osSig:
			{
				close(workers.Amount)
				fmt.Println(osSig)
				time.Sleep(time.Second * time.Duration(shutdownTimeoutSec))
				fmt.Println("Shutting down")
			}
		default:
			Gen(workers.Amount)
		}
	}

}
