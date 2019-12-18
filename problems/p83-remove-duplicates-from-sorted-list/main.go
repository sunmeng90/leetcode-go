package main

import (
	"fmt"
	model "github.com/sunmeng90/leetcode-go/model"
)

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

	Input: 1->1->2
	Output: 1->2
Example 2:

	Input: 1->1->2->3->3
	Output: 1->2->3
*/
func deleteDuplicates(head *model.ListNode) *model.ListNode {
	cur := head
	uniqueNode := head
	for cur != nil {
		if cur.Next != nil && cur.Next.Val != cur.Val {
			uniqueNode.Next = cur.Next
			uniqueNode = cur.Next
		}
		cur = cur.Next
	}
	if uniqueNode != nil && uniqueNode.Next != nil && uniqueNode.Val == uniqueNode.Next.Val { // handle empty list and  remove duplicate from end
		uniqueNode.Next = nil
	}
	return head
}

func main() {
	head := &model.ListNode{
		Val:  0,
		Next: nil,
	}
	n1 := &model.ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = n1
	n2 := &model.ListNode{
		Val:  1,
		Next: nil,
	}
	n1.Next = n2
	n3 := &model.ListNode{
		Val:  2,
		Next: nil,
	}
	n2.Next = n3
	n4 := &model.ListNode{
		Val:  2,
		Next: nil,
	}
	n3.Next = n4
	n5 := &model.ListNode{
		Val:  3,
		Next: nil,
	}
	n4.Next = n5
	n6 := &model.ListNode{
		Val:  4,
		Next: nil,
	}
	n5.Next = n6
	n7 := &model.ListNode{
		Val:  4,
		Next: nil,
	}
	n6.Next = n7
	fmt.Printf("%s\n", head.ToString())
	deleteDuplicates(head)
	fmt.Printf("%s\n", head.ToString())
	fmt.Printf("%s\n", deleteDuplicates(nil).ToString())
}
