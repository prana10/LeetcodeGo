package main

import (
	"fmt"
	"leetcode-go/leetcode"
)

func main() {
	fmt.Println("Hello, World Leetcode!")
	// Contoh input pertama: l1 = [2,4,3], l2 = [5,6,4]
	l1 := &leetcode.ListNode{Val: 2, Next: &leetcode.ListNode{Val: 4, Next: &leetcode.ListNode{Val: 3}}}
	l2 := &leetcode.ListNode{Val: 5, Next: &leetcode.ListNode{Val: 6, Next: &leetcode.ListNode{Val: 4}}}

	result := leetcode.AddTwoNumbers(l1, l2)

	// Cetak hasil
	fmt.Print("Hasil: [")
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
		if result != nil {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}
