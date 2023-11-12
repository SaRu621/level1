package main

import (
	"fmt"
	"sync"
)

const n = 5

func calcSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done() //wg.Done() выполнится перед закрытием функции calcSquare, wg.Done() увеличивает счетчик отработавших горутин для прохождения точки wg.Wait()

	fmt.Println(num * num)
}

func main() {
	sl := make([]int, n) //срез длины sl

	for i := range sl {
		sl[i] = (i + 1) * 2
	}

	fmt.Println(sl)

	var wg sync.WaitGroup
	wg.Add(n) //добавляем n горутин в wait group

	for i := range sl {
		go calcSquare(sl[i], &wg) // запуск функции calcSquare в горутинах
	}

	wg.Wait() //ждем добавленные горутины в wait group и идем дальше
}
