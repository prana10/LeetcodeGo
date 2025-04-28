package main

import (
	"fmt"
	"leetcode-go/leetcode"
)

func main() {
	fmt.Println("Hello, World Leetcode!")

	nums1 := []int{1, 3}
	nums2 := []int{2}

	fmt.Println(leetcode.FindMedianSortedArrays(nums1, nums2))
}
