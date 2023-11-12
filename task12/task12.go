package main

import (
	"fmt"
)

func makeSet(sl []string) map[string]bool {
	mp := make(map[string]bool)

	for i := range sl { //добавляем в map (множество)
		if _, exists := mp[sl[i]]; !exists {
			mp[sl[i]] = false
		}
	}

	return mp
}

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(sl)
	fmt.Println(makeSet(sl))
}
