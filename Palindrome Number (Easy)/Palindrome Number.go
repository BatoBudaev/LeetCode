package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	n := len(s)

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	x := 121
	res := isPalindrome(x)

	fmt.Println(res)
}
