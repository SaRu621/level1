package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 20

func Graduation(temp []float64) map[int][]float64 {
	mp := make(map[int][]float64)

	for i := range temp {
		key := int((temp[i] / 10)) * 10 //выделяем количество десятков градусов для определения ключа

		if _, exists := mp[key]; exists { //если группа уже существует, то:
			sl := append(mp[key], temp[i])
			mp[key] = sl
		} else { //иначе:
			sl := make([]float64, 1)
			sl[0] = temp[i]
			mp[key] = sl
		}
	}

	return mp
}

func main() {
	rand.Seed(time.Now().UnixNano()) //псевдослучайные числа генерируются по-разному каждый запуск программы

	sl := make([]float64, n)

	for i := range sl {
		sl[i] = float64(rand.Intn(801)-400) / 10.0
	}

	fmt.Println(sl)

	mp := Graduation(sl)

	fmt.Println(mp)
}
