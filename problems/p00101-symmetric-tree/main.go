package main

import (
	list "container/list"
	"github.com/sunmeng90/leetcode-go/model"
)

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Note:
Bonus points if you could solve it both recursively and iteratively.
*/
func isSymmetric(root *model.TreeNode) bool {
	nodes := make([]*model.TreeNode, 0)
	InOrder(root, func(node *model.TreeNode) {
		nodes = append(nodes, node)
	})
	if len(nodes) == 0 {
		return true
	}
	if len(nodes)%2 == 0 {
		return false
	}
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		if !equal(nodes[i], nodes[j]) || !equal(nodes[i].Left, nodes[j].Right) || !equal(nodes[i].Right, nodes[j].Left) {
			return false
		}
	}
	return true
}

func InOrder(root *model.TreeNode, nodeFun func(node *model.TreeNode)) {
	if root == nil {
		return
	}
	InOrder(root.Left, nodeFun)
	nodeFun(root)
	InOrder(root.Right, nodeFun)
}

func equal(m *model.TreeNode, n *model.TreeNode) bool {
	if m == nil && n == nil {
		return true
	}
	if m != nil && n != nil && m.Val == n.Val {
		return true
	}
	return false
}

//check node are mirror, no extra space needed
func isSymmetric2(root *model.TreeNode) bool {
	return isMirror(root, root)
}

/**
two node are mirror if their values equals and left, right children of one are mirror of another right, left
*/
func isMirror(t1 *model.TreeNode, t2 *model.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

//iterative way
func isSymmetric3(root *model.TreeNode) bool {
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(root)
	for queue.Len() > 0 {
		t1 := queue.Remove(queue.Front()).(*model.TreeNode)
		t2 := queue.Remove(queue.Front()).(*model.TreeNode)
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		queue.PushBack(t1.Left)
		queue.PushBack(t2.Right)
		queue.PushBack(t1.Right)
		queue.PushBack(t2.Left)
	}
	return true
}

func main() {
	root := &model.TreeNode{Val: 1}
	isSymmetric3(root)
}
