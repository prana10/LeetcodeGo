package main

import (
	"fmt"

	"leetcode-go/leetcode"
)

func main() {
	fmt.Println("Hello, World Leetcode!")

	example1 := "abcabcbb"
	example2 := "bbbbb"
	example3 := "pwwkew"

	fmt.Println(leetcode.LengthOfLongestSubstring(example1))
	fmt.Println(leetcode.LengthOfLongestSubstring(example2))
	fmt.Println(leetcode.LengthOfLongestSubstring(example3))
}
