package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(2 << 11)
	y := big.NewInt(2 << 10)
	fmt.Println(x, y)
	z := new(big.Int)

	z.Mul(x, y) //умножение
	fmt.Println("x * y = ", z)

	z.Div(x, y) //деление
	fmt.Println("x / y = ", z)

	z.Add(x, y) //сложение
	fmt.Println("x + y = ", z)

	z.Sub(x, y) //вычитаение
	fmt.Println("x - y = ", z)
}
