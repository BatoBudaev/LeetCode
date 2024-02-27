package main

import "fmt"

func twoSum(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			if i == j {
				break
			}

			if nums[i]+nums[j] == target {
				output := []int{i, j}
				return output
			}
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 10

	output := twoSum(nums, target)
	fmt.Println(output)
}
