package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Must provide a flags")
		panic("Must provide a flags")
	}
	filename := os.Args[1]
	message := os.Args[2]

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		log.Fatal(err)
	}

}

func flagSample() {
	file := flag.String("file", "test.txt", "Read File")
	f, err := os.OpenFile(*file, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	//text := []byte("\n Testing")
	_, err = f.WriteString("\n Testing")
	if err != nil {
		log.Fatal(err)
		return
	}
}
