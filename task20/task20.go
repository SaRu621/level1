package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func StringGenerator(n int) string {
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

func PhraseGenerator(n, m int) string { //генерируем фразу
	phrase := ""

	for i := 0; i < n; i++ {
		if i != n-1 {
			phrase += StringGenerator(rand.Intn(m)+1) + " "
		} else {
			phrase += StringGenerator(rand.Intn(m) + 1)
		}
	}

	return phrase
}

func ReversePhrase(phrase string) string {
	phraseSlice := strings.Split(phrase, " ") //дробим строку и получаем []string

	for i := 0; i < len(phraseSlice)/2; i++ {
		phraseSlice[i], phraseSlice[len(phraseSlice)-1-i] = phraseSlice[len(phraseSlice)-1-i], phraseSlice[i] //меняем зеркально местами
	}

	return strings.Join(phraseSlice, " ")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	phrase := PhraseGenerator(4, 10)
	fmt.Println(phrase)
	fmt.Println(ReversePhrase(phrase))
}
