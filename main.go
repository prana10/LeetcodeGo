package main

import (
	"fmt"
	"leetcode-go/leetcode"
)

func main() {
	fmt.Println("Hello, World Leetcode!")

	nums := []int{2, 7, 11, 15}
	target := 9

	result := leetcode.TwoSum(nums, target)
	fmt.Println("result: ", result)
}
