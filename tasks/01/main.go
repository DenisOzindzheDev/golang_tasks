package main

import "fmt"

type Human struct {
	Id   int
	Name string
}
type Action struct {
	Human
}

func (h *Human) Greetings() string {
	return fmt.Sprintf("NAME %s", h.Name)
}

func main() {

	myAction := Action{Human{Id: 0, Name: "Denzel"}}
	fmt.Print(myAction.Greetings())
	myAction.Name = "Gnoverk"
	fmt.Print(myAction.Greetings())
}
