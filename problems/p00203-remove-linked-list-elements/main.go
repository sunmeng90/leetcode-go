package main

import "github.com/sunmeng90/leetcode-go/model"

/*
Remove all elements from a linked list of integers that have value val.

Example:

	Input:  1->2->6->3->4->5->6, val = 6
	Output: 1->2->3->4->5
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *model.ListNode, val int) *model.ListNode {
	result := &model.ListNode{}
	resultHead := result
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val != val {
			result.Next = cur
			result = cur
		} else if cur.Next == nil {
			result.Next = nil
		}
	}
	return resultHead.Next
}
