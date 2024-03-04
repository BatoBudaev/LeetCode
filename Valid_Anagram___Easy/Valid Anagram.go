package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		map1[rune(s[i])] += 1
		map2[rune(t[i])] += 1
	}

	return reflect.DeepEqual(map1, map2)
}

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}
