package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	sort.Ints(nums3)

	mid := len(nums3) / 2

	if len(nums3)%2 == 0 {
		return float64(nums3[mid-1]+nums3[mid]) / 2.0
	}

	return float64(nums3[mid])
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	res := findMedianSortedArrays(nums1, nums2)

	fmt.Println(res)
}
