package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Test string `json:"test,omitempty"`
	Mama string `json:"mama,omitempty"`
}

func main() {
	c := Config{
		Test: "test",
		Mama: "mama",
	}
	test, _ := json.Marshal(c) //serialize
	fmt.Printf("aa %v", json.Unmarshal(test, &c))
}
