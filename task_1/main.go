package main

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
import "fmt"

type Human struct {
	name string
	age  int
	sex  string
	job  string
}
type Animal struct {
	name string
	age  int
}

func (h *Human) Run() {
	fmt.Printf("%s is running\n", h.name)
}
func (a *Animal) Run() {
	fmt.Printf("%s is running\n", a.name)
}

type Action struct {
	Human
	Animal
}

func (a *Action) Run() {
	a.Animal.Run()
	fmt.Println("Action is running")
}
func main() {
	human := Human{
		name: "Denis",
		age:  25,
		sex:  "male",
		job:  "programmer",
	}
	cat := Animal{
		name: "Cat",
		age:  2,
	}
	action := Action{
		Human:  human,
		Animal: cat,
	}
	cat.Run()
	human.Run()
	action.Run()

}
