package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func StringGenerator(n int) string { //генерируем строку
	str := ""

	for i := 0; i < n; i++ {
		flag := rand.Intn(2)

		if flag == 0 {
			str += string(rune((65 + rand.Intn(26))))
		} else {
			str += strconv.Itoa(rand.Intn(10))
		}
	}

	return str
}

func reverseString(str string) string {
	runeStr := []rune(str) //переводим в []rune для корректного считывания символов

	for i := 0; i < len(runeStr)/2; i++ { //идем от начала к середине
		runeStr[i], runeStr[len(runeStr)-i-1] = runeStr[len(runeStr)-i-1], runeStr[i] //меняем местами зеркально
	}

	return string(runeStr) //обратно в []byte
}

func main() {
	rand.Seed(time.Now().UnixNano())
	str := StringGenerator(5)

	fmt.Println(str)

	str = reverseString(str)

	fmt.Println(str)
}
