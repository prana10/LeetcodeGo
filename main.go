package main

import (
	"fmt"
	"leetcode-go/leetcode"
)

func main() {
	fmt.Println("Hello, World Leetcode!")
	// minK, maxK (untuk membatasi nilai min & max dari array)
	minK := 1
	maxK := 1
	fmt.Println("minK: ", minK)
	fmt.Println("maxK: ", maxK)

	// array of integers
	nums := []int{1, 1, 1, 1}
	fmt.Println("length of nums: ", len(nums))

	// count subarrays
	result := leetcode.CountSubarrays(nums, minK, maxK)
	fmt.Println("result: ", result)
}
