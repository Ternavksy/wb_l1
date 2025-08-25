package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) SayHello() {
	fmt.Printf("Привет, я %s, мне %d года.\n", h.Name, h.Age)
}

type Action struct {
	Human
	Activity string
}

func (a Action) DoSomething() {
	fmt.Printf("Что делает? %s.\n", a.Activity)
}

func main() {
	newHuman := Action{
		Human: Human{
			Name: "Даня",
			Age:  22,
		},
		Activity: "Выполняет курс техношколы WB",
	}
	newHuman.SayHello()
	newHuman.DoSomething()
}
