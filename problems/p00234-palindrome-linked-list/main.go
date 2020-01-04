package main

import (
	"github.com/sunmeng90/leetcode-go/model"
	"github.com/sunmeng90/leetcode-go/util"
)

/*
Given a singly linked list, determine if it is a palindrome.

Example 1:
	Input: 1->2
	Output: false

Example 2:
	Input: 1->2->2->1
	Output: true
Follow up:
Could you do it in O(n) time and O(1) space?
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *model.ListNode) bool {
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}
	if count < 2 {
		return true
	}
	mid := count / 2 // the latter one
	firstHalf := make([]*model.ListNode, 0)
	idx, idx2 := 0, 0
	for cur := head; cur != nil; cur = cur.Next {
		if idx < mid {
			firstHalf = append(firstHalf, cur)
			idx2 = idx + 1
		} else if idx > mid {
			idx2--
			if firstHalf[idx2].Val != cur.Val {
				return false
			}
		} else if idx == mid {
			if count%2 == 0 {
				idx2--
				if firstHalf[idx2].Val != cur.Val {
					return false
				}
			}
		}
		idx++
	}
	return true
}

func main() {
	data := util.SliceToSinglyLinkedList([]int{1, 0, 0})
	println(data.ToString())
	println(isPalindrome(data))
}
