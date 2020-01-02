package main

import "github.com/sunmeng90/leetcode-go/model"

/*
Invert a binary tree.

Example:

Input:
		 4
	   /   \
	  2     7
	 / \   / \
	1   3 6   9

Output:
		 4
	   /   \
	  7     2
	 / \   / \
	9   6 3   1

Trivia:
This problem was inspired by this original tweet by Max Howell:

Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *model.TreeNode) *model.TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
