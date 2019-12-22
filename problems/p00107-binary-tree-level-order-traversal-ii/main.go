package main

import (
	"container/list"
	"fmt"
	"github.com/sunmeng90/leetcode-go/model"
)

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
		3
	   / \
	  9  20
		/  \
	   15   7
return its bottom-up level order traversal as:
	[
	  [15,7],
	  [9,20],
	  [3]
	]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *model.TreeNode) [][]int {
	queue := list.New()
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue.PushBack(root) // interface value added
	for queue.Len() > 0 {
		nodesCountInLevel := queue.Len()
		nodes := make([]int, 0)
		for i := 0; i < nodesCountInLevel; i++ { // add children for all nodes in current level
			node := queue.Remove(queue.Front()).(*model.TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			nodes = append(nodes, node.Val) // after add its children, then add current node to the result
		}
		result = append(result, nodes)
	}
	reverseSlice(result)
	return result
}

func reverseSlice(data [][]int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func levelOrderBottom2(root *model.TreeNode) [][]int {
	result := make([]*[]int, 0)
	doPopulateLevel(&result, root, 0)
	height := len(result)
	deRefResult := make([][]int, height)
	for i := height - 1; i >= 0; i-- {
		deRefResult[height-i-1] = *result[i]
	}
	return deRefResult
}

func doPopulateLevel(result *[]*[]int, root *model.TreeNode, level int) {
	if root == nil {
		return
	}
	if level >= len(*result) { //new level
		nodesInLevel := make([]int, 0)
		*result = append(*result, &nodesInLevel) // result change
	}
	doPopulateLevel(result, root.Left, level+1)
	doPopulateLevel(result, root.Right, level+1)
	*(*result)[level] = append(*(*result)[level], root.Val) // add current node to the sub-list
}

func main() {
	root := &model.TreeNode{Val: 3}
	n9 := &model.TreeNode{Val: 9}
	n20 := &model.TreeNode{Val: 20}
	n15 := &model.TreeNode{Val: 15}
	n7 := &model.TreeNode{Val: 7}
	root.Left = n9
	root.Right = n20
	n20.Left = n15
	n20.Right = n7
	fmt.Printf("%v\n", levelOrderBottom2(root))
}
