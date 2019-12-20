package main

import (
	"github.com/sunmeng90/leetcode-go/model"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *model.TreeNode, q *model.TreeNode) bool {
	left, right := p, q
	if left == right {
		return true
	}
	if left != nil && right != nil && left.Val == right.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
