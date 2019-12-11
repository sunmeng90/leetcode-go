package main

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

	Input: 1->2->4, 1->3->4
	Output: 1->1->2->3->4->4
*/

/*
Solution:
	two pointer: head(result with a fake empty value first), cur(move on the result when iteration)
	find the min value from the two list, and append the min value to the result until the min value is nil (no elements)
	we need move head once to remove the fake empty value

*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{} // a fake empty value
	cur := head
	for {
		// get the min value from two list
		var rtn *ListNode
		if l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				rtn = l2
				l2 = l2.Next
			} else {
				rtn = l1
				l1 = l1.Next
			}
		} else if l1 != nil {
			rtn = l1
			l1 = l1.Next
		} else {
			rtn = l2
			if l2 != nil {
				l2 = l2.Next
			}
		}
		// append to cur with min value
		if rtn != nil {
			cur.Next = rtn
			cur = cur.Next
		} else {
			break
		}
	}
	return head.Next
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
