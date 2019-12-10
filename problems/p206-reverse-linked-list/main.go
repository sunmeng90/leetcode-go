package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func reverseList(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var pre, next *listNode
	for cur.Next != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	cur.Next = pre
	return cur
}

func main() {
	head := &listNode{Val: 1}
	var cur = head
	for i := 2; i < 6; i++ {
		cur.Next = &listNode{Val: i}
		cur = cur.Next
	}
	reversed := reverseList(head)
	for v := reversed; v != nil; v = v.Next {
		fmt.Println(v.Val)
	}
}
