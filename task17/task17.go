package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func binarySearch(sl []int, toFind int) int {
	leftBorder, rightBorder := 0, len(sl) //левая граница указывает на начало отрезка, правая -- на элемент за последним элементом
	for rightBorder-leftBorder > 1 {      //пока отрезок больше единицы
		middle := (leftBorder + rightBorder) / 2 //середина отрезка

		if sl[middle] < toFind { //если меньше, то пододвигаем левую границу
			leftBorder = middle
		} else if sl[middle] > toFind { //иначе -- правую
			rightBorder = middle
		} else {
			return middle
		}
	}

	return -1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	sl := make([]int, n)

	for i := range sl {
		if i == 0 {
			sl[i] = 1
		} else {
			sl[i] = rand.Intn(3) + 1 + sl[i-1]
		}
	}

	fmt.Println(sl)

	fmt.Println(binarySearch(sl, 10))
}
