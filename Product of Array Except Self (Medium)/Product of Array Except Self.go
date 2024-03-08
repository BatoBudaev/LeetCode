package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftArr := make([]int, n)
	rightArr := make([]int, n)
	answer := make([]int, n)

	leftProduct := 1
	rightProduct := 1

	for i := 0; i < n; i++ {
		leftArr[i] = leftProduct
		rightArr[n-1-i] = rightProduct
		leftProduct *= nums[i]
		rightProduct *= nums[n-1-i]
	}

	for i := range leftArr {
		answer[i] = leftArr[i] * rightArr[i]
	}

	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println(productExceptSelf(nums))
}
