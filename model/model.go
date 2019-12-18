package model

import "strconv"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToString() string {
	if l == nil {
		return ""
	}
	cur := l
	result := strconv.Itoa(cur.Val)
	for cur.Next != nil {
		result = result + ", " + strconv.Itoa(cur.Next.Val)
		cur = cur.Next
	}
	return result
}
