package main

import (
	"fmt"
)

const n = 10

func main() {
	sl := make([]int, n)

	for i := range sl {
		sl[i] = i
	}

	fmt.Println(sl)

	//sl -> ch1 (->)^2 ch2
	ch1 := make(chan int) //буфер между ch2 и sl
	ch2 := make(chan int)

	go func() { //в данном замыкании добавляем значения из sl в ch1.  Замыкание в горутине. После отсылки всех элементов канал закрывается.
		defer close(ch1)

		for i := range sl {
			ch1 <- sl[i]
		}
	}()

	go func() { //в этом замыкании проходимся по ch1 и записываем квадрат в ch2. После закрываем канал.
		defer close(ch2)

		for val := range ch1 {
			ch2 <- val * val
		}
	}()

	for val := range ch2 {
		fmt.Println(val)
	}

}
