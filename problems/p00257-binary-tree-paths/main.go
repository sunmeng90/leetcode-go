package main

import (
	"github.com/sunmeng90/leetcode-go/model"
	"strconv"
)

/*
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *model.TreeNode) []string {
	result := make([]string, 0)
	if root != nil {
		concatPath(root, "", &result)
	}
	return result
}

func concatPath(root *model.TreeNode, path string, result *[]string) {
	curStr := path + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*result = append(*result, curStr)
	}
	if root.Left != nil {
		concatPath(root.Left, curStr+"->", result)
	}
	if root.Right != nil {
		concatPath(root.Right, curStr+"->", result)
	}
}

func main() {
}
