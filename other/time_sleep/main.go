package main

import (
	"log"
	"time"
)

func main() {
	timeSleep(5 * time.Second)
	log.Print("A? A? A?")
	timeSleep(1 * time.Second)
	log.Print("A? A? A?")
}

func timeSleep(ns time.Duration) {
	if ns <= 0 {
		return // Если таймер установлен в 0 или меньше 0 - то выходим из функции
	}
	<-time.After(ns)
}
