package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	return len(m) != len(nums)
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(arr))
}
