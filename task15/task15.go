package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

var justString string

var m runtime.MemStats

const toCopy = 1000000

func createHugeString(n int) string {
	str := make([]byte, n)

	for i := range str {
		str[i] = byte(65 + rand.Intn(26))
	}

	return string(str)
}

func someFunc() {
	v := createHugeString(1 << 30)

	justString = v[:toCopy] //*

	runtime.ReadMemStats(&m)

	fmt.Println(m.HeapAlloc) //HeapAlloc = 2147559720
}

func mySomeFunc() {
	v := createHugeString(1 << 30)

	vCopy := make([]byte, toCopy)

	copy(vCopy, v[:toCopy]) //**

	runtime.ReadMemStats(&m)

	fmt.Println(m.HeapAlloc) //HeapAlloc = 1073924600
}

// В чем преимущество?
// 1. В * может быть ошибка, если размер массива меньше чем toCopy
// 2. В * происходит лишнее выделение памяти, в отличии от **

func main() {
	//someFunc()
	mySomeFunc()
}

//не сделано
