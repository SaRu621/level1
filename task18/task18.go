package main

import (
	"fmt"
	"sync"
)

const n = 100000

type Counter struct {
	i int
	m sync.Mutex
}

func (c *Counter) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	c.m.Lock()
	c.i++
	c.m.Unlock()
}

func (c Counter) getIndex() int {
	return c.i
}

func main() {
	var obj Counter

	var wg sync.WaitGroup

	wg.Add(n)

	fmt.Println(obj.getIndex())

	for i := 0; i < n; i++ {
		obj.inc(&wg)
	}

	wg.Wait()

	fmt.Println(obj.getIndex())
}
