package main

import (
	"fmt"
	"regexp"
)

func main() {
	name := "Kubernetes Lohovich23"
	namePattern := `^[a-zA-Z]+(?: [a-zA-Z]+)?$`

	matchName, _ := regexp.MatchString(namePattern, name)

	fmt.Printf("Name  '%s' matching state: %v", name, matchName)

}
