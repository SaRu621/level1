package main

import (
	"fmt"
	"os"
	"time"
)

const n = 1

func recv(ch chan int) { //получение значений int из канала в stdout
	for {
		fmt.Println(<-ch)
	}
}

func send(ch chan int) { //отправка значений int в канал
	for i := 0; true; i++ {
		ch <- i
	}
}

func main() {
	fmt.Println()

	ch := make(chan int)

	go recv(ch) //в одной горутине -- получение

	go send(ch) //в другой -- отправка

	time.Sleep(time.Second * n)
	os.Exit(0)
}
