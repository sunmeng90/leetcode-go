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

func binaryTreePathsNonRecursive(root *model.TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}
	nodes := make([]*model.TreeNode, 0)
	nodes = append(nodes, root)
	tmpPath := make([]string, 0)
	tmpPath = append(tmpPath, strconv.Itoa(root.Val))
	for len(nodes) != 0 {
		cur := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		curPath := tmpPath[len(tmpPath)-1]
		tmpPath = tmpPath[:len(tmpPath)-1]
		if cur.Left == nil && cur.Right == nil {
			result = append(result, curPath)
			continue
		}
		if cur.Left != nil {
			nodes = append(nodes, cur.Left)
			tmpPath = append(tmpPath, curPath+"->"+strconv.Itoa(cur.Left.Val))
		}
		if cur.Right != nil {
			nodes = append(nodes, cur.Right)
			tmpPath = append(tmpPath, curPath+"->"+strconv.Itoa(cur.Right.Val))
		}
	}
	return result
}

func main() {
	nodes := make([]*model.TreeNode, 0)
	println(nodes[:0])
	root := &model.TreeNode{Val: 1}
	n2 := &model.TreeNode{Val: 2}
	n3 := &model.TreeNode{Val: 3}
	n5 := &model.TreeNode{Val: 5}
	root.Left = n2
	root.Right = n3
	n2.Right = n5
	println(binaryTreePathsNonRecursive(root))

}
