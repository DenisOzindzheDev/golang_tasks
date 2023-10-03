```
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
	fmt.Println("Action is running")
}

```

[Collision](https://golangify.com/composition-and-forwarding )

Композиция в GO 
структуры и подструктуры