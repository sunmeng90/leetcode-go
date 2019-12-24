package main

import (
	"github.com/sunmeng90/leetcode-go/model"
	"github.com/sunmeng90/leetcode-go/util"
)

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the left and right subtrees of every node differ in height by no more than 1.



Example 1:

Given the following tree [3,9,20,null,null,15,7]:

		3
	   / \
	  9  20
		/  \
	   15   7
	Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

		   1
		  / \
		 2   2
		/ \
	   3   3
	  / \
	 4   4
Return false.
*/

// Tree is balanced if left and right max dept difference is less than 1 and both of them are balanced
func isBalanced(root *model.TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return util.Abs(maxDepth(root.Left)-maxDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *model.TreeNode) int {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			return 1
		}
		if root.Left != nil && root.Right != nil {
			return 1 + util.Max(maxDepth(root.Left), maxDepth(root.Right))
		}
		if root.Left != nil {
			return 1 + maxDepth(root.Left)
		}
		if root.Right != nil {
			return 1 + maxDepth(root.Right)
		}
	}
	return 0
}
