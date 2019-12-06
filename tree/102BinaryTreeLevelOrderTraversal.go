package tree

import (
	"leetcode/dataStructure"
	"leetcode/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
Given binary tree [3,9,20,null,null,15,7],
[
  [3],
  [9,20],
  [15,7]
]
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	rootLevel := make([]int, 0)
	rootLevel = append(rootLevel, root.Val)
	res = append(res, rootLevel)

	nodes := make([]*TreeNode, 0)
	if root.Left != nil {
		nodes = append(nodes, root.Left)
	}
	if root.Right != nil {
		nodes = append(nodes, root.Right)
	}
	return inner102(nodes, res)
}

func inner102(nodes []*TreeNode, res [][]int) [][]int {
	if len(nodes) == 0 {
		return res
	}
	outNodes := make([]*TreeNode, 0)
	curLevel := make([]int, 0)
	for _, root := range nodes {
		if root.Left != nil {
			outNodes = append(outNodes, root.Left)
		}
		if root.Right != nil {
			outNodes = append(outNodes, root.Right)
		}
		curLevel = append(curLevel, root.Val)
	}
	res = append(res, curLevel)
	return inner102(outNodes, res)
}

func Test102() {
	tree := []string{"1", "2", "3"}
	root := dataStructure.BuildTree(tree)
	res := levelOrder(root)
	up := utils.PrintUtil{}
	up.Print2DInt(res)
}
