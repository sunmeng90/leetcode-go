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
	return sumOfLeftLeavesOfParent(root.Left, root) + sumOfLeftLeavesOfParent(root.Right, root)
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

func sumOfLeftLeaves2(root *model.TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	stack := make([]*model.TreeNode, 0)
	stack = append(stack, root)
	sum := 0
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				sum += node.Left.Val
			} else {
				stack = append(stack, node.Left)
			}
		}
		if node.Right != nil {
			if node.Right.Left != nil || node.Right.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return sum
}

func main() {

}
