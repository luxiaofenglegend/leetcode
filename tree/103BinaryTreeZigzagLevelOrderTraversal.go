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
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := make([]*TreeNode, 0, 1)
	nodes = append(nodes, root)
	curRes := make([][]int, 0)
	return inner103(nodes, curRes, true)
}

func inner103(nodes []*TreeNode, curRes [][]int, isLeftToRight bool) [][]int {
	length := len(nodes)
	if length == 0 {
		return curRes
	}
	outNodes := make([]*TreeNode, 0, 2*length)
	curLevel := make([]int, 0, length)
	for _, root := range nodes {
		if root.Left != nil {
			outNodes = append(outNodes, root.Left)
		}
		if root.Right != nil {
			outNodes = append(outNodes, root.Right)
		}
		curLevel = append(curLevel, root.Val)
	}
	if !isLeftToRight {
		curLevel = reverse103(curLevel)
	}
	curRes = append(curRes, curLevel)
	return inner103(outNodes, curRes, !isLeftToRight)
}

func reverse103(arr []int) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}
	i, j := 0, len(arr)-1
	for j-i > 0 {
		tmp := arr[i]
		arr[i] = arr[j]
		arr[j] = tmp
		i++
		j--
	}
	return arr
}

func Test103() {
	tree := []string{"3", "9", "20", "null", "null", "15", "7"}
	root := dataStructure.BuildTree(tree)
	res := zigzagLevelOrder(root)
	up := utils.PrintUtil{}
	up.Print2DInt(res)
}
