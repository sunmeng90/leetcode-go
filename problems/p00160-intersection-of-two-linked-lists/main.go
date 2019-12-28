package main

import "github.com/sunmeng90/leetcode-go/model"

/*
Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:


begin to intersect at node c1.


Example 1:
	Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
	Output: Reference of the node with value = 8
Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.


Example 2:


	Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
	Output: Reference of the node with value = 2
Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.


Example 3:


	Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
	Output: null
Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.


Notes:

If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
if two list has the same length, they will reach the intersection point at the same time.

if not the same length, the shorter one will first go to the end, and then redirect to the head of longer list,
and the longer one will continue to finish the remaining elements, and reach the end of longer list then redirects
to the head of short list. so they will eventually meet at the intersection point

The redirect will make the remaining elements for pointer on longer one and shorter one has the same length to go
*/
func getIntersectionNode(headA, headB *model.ListNode) *model.ListNode {
	na, nb := headA, headB
	for ; na != nil && nb != nil; na, nb = na.Next, nb.Next {
		if na == nb { // same length,should meet at the interception point
			return na
		}
	}

	if na == nil && nb == nil { // not the same length, shorter one goes to end first
		return nil
	}

	if na == nil { // list a is shorter
		// redirect na to list b
		for na = headB; na != nil && nb != nil; na, nb = na.Next, nb.Next {
		}
		// nb goes to the end of b now
		for nb = headA; na != nil && nb != nil; na, nb = na.Next, nb.Next {
			if na == nb {
				return na
			}
		}

	} else if nb == nil {
		// redirect nb to list a
		for nb = headA; na != nil && nb != nil; na, nb = na.Next, nb.Next {
		}
		// na goes to the end of a now
		for na = headB; na != nil && nb != nil; na, nb = na.Next, nb.Next {
			if na == nb {
				return na
			}
		}
	}
	return nil
}

// same, but less code
func getIntersectionNode2(headA, headB *model.ListNode) *model.ListNode {
	na, nb := headA, headB
	for na != nb {
		if na == nil {
			na = headB
		} else {
			na = na.Next
		}
		if nb == nil {
			nb = headA
		} else {
			nb = nb.Next
		}
	}
	return na
}

// TODO: use map: put elements in one list and then traverse the second list to check whether elements are already in the map

// TODO: if two list have intersection, the should have the same end element, so traverse the two list and check the
// equality of the end elements, but this approach can't find the intersection element
