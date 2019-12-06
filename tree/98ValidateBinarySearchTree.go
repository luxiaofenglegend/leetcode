package tree

import (
	"fmt"
	"leetcode/dataStructure"
)

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	val := root.Val
	if root.Left != nil && root.Left.Val >= val {
		return false
	}
	if root.Right != nil && root.Right.Val <= val {
		return false
	}
	return inner98(root.Left, -9223372036854775808, val) && inner98(root.Right, val, 9223372036854775807)
}

func inner98(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	fmt.Println(root.Val, min, max)
	val := root.Val
	if val <= min || val >= max {
		return false
	}
	if root.Left != nil && root.Left.Val >= val {
		return false
	}
	if root.Right != nil && root.Right.Val <= val {
		return false
	}
	return inner98(root.Left, min, val) && inner98(root.Right, val, max)
}

func Test98() {
	arr := []string{"3", "1", "5", "0", "2", "4", "6", "null", "null", "null", "3"}
	root := dataStructure.BuildTree(arr)
	//root.PrintTree()
	isValidBST(root)
}
