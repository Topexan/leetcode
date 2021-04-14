package Leetcode

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0
	n := -1.0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		res += cur.Val * int(math.Pow(2.0, n))
		n--
		cur = cur.Next
	}
	return res
}
