package main

import (
	"fmt"

	"github.com/tayalone/go-algo/twopointers"
)

// nums: []int{1, 2, 3, 4, 5, 6, 7},
// 				k:    3,

func main() {
	fmt.Println(twopointers.RemoveNthFromEnd(&twopointers.ListNode{
		Val: 1,
		Next: &twopointers.ListNode{
			Val: 2,
			Next: &twopointers.ListNode{
				Val: 3,
				Next: &twopointers.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}, 1))
	// fmt.Println(1 ^ 1)

	fmt.Println("Hello Algo 101")
}
