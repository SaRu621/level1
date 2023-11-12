package main

import (
	"fmt"
)

func printType(x interface{}) { //аргумент -- любой тип, реализующий интерфейс (здесь интерфейс пустой -> то есть любой тип)
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel int")
	}
}

func main() {
	x := true
	printType(x)
}
