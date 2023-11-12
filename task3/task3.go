package main

import (
	"fmt"
	"sync"
)

const n = 3

var mtx sync.Mutex

func add(n *int, x int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	m.Lock() //блокируем множественный доступ к *n
	(*n) += x * x
	m.Unlock() //разблокируем доступ к *n для других горутин
}

func main() {
	sl := make([]int, n)

	for i := range sl {
		sl[i] = 2 * (i + 1)
	}

	fmt.Println(sl)

	var wg sync.WaitGroup
	wg.Add(n)

	sum := 0
	for i, _ := range sl {
		go add(&sum, sl[i], &mtx, &wg)
	}

	wg.Wait()

	fmt.Println(sum)

}
