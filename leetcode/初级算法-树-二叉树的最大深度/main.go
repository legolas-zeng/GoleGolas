package main

import "math"

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/////////////////  使用递归 //////////////////////
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return int(math.Max(float64(left), float64(right)) + 1)
}

////////////////////////////////////////////////////
