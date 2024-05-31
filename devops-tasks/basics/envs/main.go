package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Your system username: %s\n", os.Getenv("USER"))
}
