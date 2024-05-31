package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// path := "folder"

	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// 	err = os.Mkdir(path, 0755)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	data, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
