package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	str = strings.ToLower(str) //переводим всю строку в нижний регистр, чтобы регистр не имел разницы

	runeStr := []rune(str)

	mp := make(map[rune]bool, len(runeStr))

	for i := range runeStr {
		if _, ok := mp[runeStr[i]]; ok {
			return false
		} else {
			mp[runeStr[i]] = true
		}
	}

	return true
}

func main() {
	str := "tavkgdt"
	fmt.Println(isUnique(str))
}
