package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 100

func intersection(mp1, mp2 map[int]bool) map[int]bool {
	mp := make(map[int]bool)

	for key, _ := range mp1 { //проходим элементы первого множества и ищем их во втором, если они есть, тогда записываем в новое множество
		if _, exists := mp2[key]; exists {
			mp[key] = false
		}
	}

	return mp
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mp1 := make(map[int]bool) //первое множество
	mp2 := make(map[int]bool) //второе множество
	//Почему map[int]bool? Нам нужно множество, а не словарь -> нужно минимизировать память -> bool -- лучший выбор (1 байт). Пустая структура занимает минимум 1 байт.

	for i := 0; i <= n; i += 2 {
		mp1[rand.Intn(n)] = false
	}

	for i := 0; i <= n; i += 5 {
		mp2[rand.Intn(n)] = false
	}

	//fmt.Println(mp1, mp2)

	mp := intersection(mp1, mp2)

	for key, _ := range mp {
		fmt.Print(key, " ")
	}
	fmt.Println()
}
