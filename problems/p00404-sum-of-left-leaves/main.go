package main

import "github.com/sunmeng90/leetcode-go/model"

/*
Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *model.TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	return sumOfLeftLeavesOfParent(root.Left, root) +  sumOfLeftLeavesOfParent(root.Right, root)
}

func sumOfLeftLeavesOfParent(node *model.TreeNode, parent *model.TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil && parent.Left == node {
		return node.Val
	}
	return sumOfLeftLeavesOfParent(node.Left, node) + sumOfLeftLeavesOfParent(node.Right, node)
}

func main() {

}
