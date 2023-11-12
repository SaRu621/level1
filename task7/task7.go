package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type concurrencyMap struct {
	m  sync.Mutex
	mp map[string]int
}

func (cm *concurrencyMap) add(key string, val int, wg *sync.WaitGroup) {
	defer wg.Done()

	cm.m.Lock() //даем доступ к добавлению значения только одной горутине за раз
	defer cm.m.Unlock()
	cm.mp[key] = val
}

const n = 5

func main() {
	obj := concurrencyMap{ //Первый способ: использование  map + sync.Mutex
		mp: make(map[string]int),
	}

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go obj.add(StringGenerator(10), i, &wg) //добавление элементов в горутинах
	}

	wg.Wait()

	fmt.Println(len(obj.mp))

	var sMap sync.Map //Второй способ: использование  sync.Map

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			sMap.Store(StringGenerator(10), i) //добавление значения в sync.Map
		}(&wg)
	}

	wg.Wait()

	fmt.Println(sMap)
}

func StringGenerator(n int) string { //генератор строк
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
