package main

import (
	"github.com/sunmeng90/leetcode-go/model"
	"github.com/sunmeng90/leetcode-go/util"
)

/*
if node is nil, then depth is 0
if node is not nil, then:
	a) if left and right is both nil, then depth is 1
	b) if left and right are both non-nil, then max depth is the max value between lef and right + 1
	c) if left not nil, then max depth is max depth of left tree +1
	c) if right not nil, then max depth is max depth of right tree +1
*/
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
