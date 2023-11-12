package main

import "fmt"

func changeBitTo(num, i int, flag bool) int { //num -- число, i -- бит, flag -- если true, то меняем бит на 1 , если false, то на 0
	mask := 1 << i //i = 0: 1, i = 1: 10, i = 2: 100 ... (1, 10, 100 в двоичной системе)

	if flag {
		return (num | mask) // побитовое "или" с маской
	} else {
		return (num & ^mask) //побитовое "и" с инвертированием битов маски
	}
}

func main() {
	for i := 0; i < 7; i++ {
		fmt.Println(changeBitTo(127, i, false))
	}
}
