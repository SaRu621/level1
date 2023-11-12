package main

import (
	"fmt"
	"sync"
)

func worker(x chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker num: %d\n", <-x)
}

func main() {
	var num int
	fmt.Println("Enter num of workers:")
	fmt.Scan(&num)

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go worker(ch, &wg)
	}

	i := 0
	for {
		select {
		case ch <- i:
			i++
		default:
		}
	}

	wg.Wait()

}
