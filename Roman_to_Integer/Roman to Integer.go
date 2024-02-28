package main

import "fmt"

func romanToInt(s string) int {
	m := make(map[rune]int)

	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	res := 0
	var prevNum int

	for _, c := range s {
		num := m[c]
		res += num

		if prevNum*5 == num || prevNum*10 == num {
			res -= prevNum * 2
		}

		prevNum = num
	}

	return res
}

func main() {
	s := "MCMXCIV"

	fmt.Println(romanToInt(s))
}
