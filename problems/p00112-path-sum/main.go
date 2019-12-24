package main

import "github.com/sunmeng90/leetcode-go/model"

func hasPathSum(root *model.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		}
		return false
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

/*
[1,-2,-3,1,3,-2,null,-1]
2
true
*/
func main() {
	root := &model.TreeNode{Val: -2}
	n3 := &model.TreeNode{Val: -3}
	root.Right = n3
	println(hasPathSum(root, -5))
}
