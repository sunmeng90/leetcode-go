package main

import "github.com/sunmeng90/leetcode-go/model"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
We go through each node one by one and record each node's reference (or memory address) in a hash table. If the current
node is null, we have reached the end of the list and it must not be cyclic. If current node’s reference is in the hash
table, then return true.
*/
func hasCycle(head *model.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	visited := make(map[*model.ListNode]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		_, ok := visited[cur]
		if ok {
			return true
		} else {
			visited[cur] = 1
		}
	}
	return false
}

// TODO: image there are two cars in the lane, a slow car and speed car, if the lane is a cycle the fast car will eventually
// catch up the slow car somewhere. If not a cycle, the two cars will never meet
