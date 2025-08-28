package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayHello() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", h.Name, h.Age)
}

func (h *Human) GetAge() int {
	return h.Age
}

func (h *Human) SetAge(age int) {
	if age > 0 {
		h.Age = age
	}
}

type Action struct {
	*Human
	ActionName string
}

func (a *Action) DoAction() string {
	return fmt.Sprintf("%s (age %d) is doing %s", a.Name, a.Age, a.ActionName)
}

func main() {
	human := &Human{
		Name: "Alice",
		Age:  30,
	}

	action := &Action{
		Human:      human,
		ActionName: "running",
	}

	fmt.Println(action.SayHello())
	fmt.Println(action.GetAge())

	fmt.Println(action.DoAction())

	action.SetAge(31)
	fmt.Println(action.GetAge())

	fmt.Println(human.GetAge())
}
