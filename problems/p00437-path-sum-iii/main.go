package main

import (
	"fmt"
	"github.com/sunmeng90/leetcode-go/model"
)

/*
You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

Example:

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *model.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0
	var dfs func(*model.TreeNode, int)
	dfs = func(node *model.TreeNode, sum int) {
		if node == nil {
			return
		}
		sum -= node.Val
		if sum == 0 {
			cnt += 1
			// at first thought, we should return here, we already got the matched sum, and should exit, passing sum of 0
			// to sub-tree is wrong, but look at this case(sum 2): 1->1->0->0, the path set should not only include:
			// 1->1, but also 1->1->0, 1->1->0->0, that is for the matched path, if all subsequent node sums up to 0,
			// then the first match path + the subsequent path should also a matched path
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, sum)
	return cnt + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func main() {
	root := &model.TreeNode{
		Val: 1,
		Left: &model.TreeNode{
			Val: -2,
			Left: &model.TreeNode{
				Val:   1,
				Left:  &model.TreeNode{Val: -1},
				Right: nil,
			},
			Right: &model.TreeNode{Val: 3},
		},
		Right: &model.TreeNode{
			Val:   -3,
			Left:  &model.TreeNode{Val: -2},
			Right: nil,
		},
	}
	fmt.Printf("Count of sum: %d\n", pathSum(root, -1))
}
