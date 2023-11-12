package main

import (
	"fmt"
)

type Human struct {
}

func (h Human) Hello() {
	fmt.Println("Hi!")
}

type Action struct {
	Human // встраиваем структуру Human в Action (включая методы)
}

func (a Action) Bye() {
	fmt.Println("Bye!")
}

func main() {
	var t Action
	t.Hello() //метод "наследован" от структуры Human
	t.Bye()   //метод структуры action
}
