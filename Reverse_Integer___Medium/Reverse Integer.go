package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var res int

	for x/10 != 0 {
		y := x % 10
		res = res*10 + y
		x /= 10
	}

	res = res*10 + x

	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}

func main() {
	fmt.Println(reverse(120))
}
