package main

import (
	"github.com/sunmeng90/leetcode-go/model"
	"github.com/sunmeng90/leetcode-go/util"
)

/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
return its minimum depth = 2.

*/
func minDepth(root *model.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil { // leaf: return height of leaf itself
		return 1
	}
	if root.Left != nil && root.Right == nil { // root has one child
		return 1 + minDepth(root.Left)
	}
	if root.Right != nil && root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	return 1 + util.Min(minDepth(root.Left), minDepth(root.Right)) // root has two children
}

func main() {
	root := &model.TreeNode{Val: 1}
	n2 := &model.TreeNode{Val: 2}
	n3 := &model.TreeNode{Val: 3}
	n4 := &model.TreeNode{Val: 4}
	n5 := &model.TreeNode{Val: 5}
	root.Left = n2
	root.Right = n3
	n2.Left = n4
	n2.Right = n5
	println(minDepth(root))
}
