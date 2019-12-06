package dataStructure

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		Val:   -8848,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
}

func BuildTree(nodeList []string) *TreeNode {
	root := NewTreeNode()
	q := NewQueue()
	q.Push(root)
	index := 0
	for index < len(nodeList) {
		curNode := q.Top().(*TreeNode)
		q.Pop()
		//_ = curNode.Val
		//
		val := -8848
		if nodeList[index] != "null" {
			val, _ = strconv.Atoi(nodeList[index])
		}
		curNode.Val = val
		index++
		curNode.Left = NewTreeNode()
		curNode.Right = NewTreeNode()
		q.Push(curNode.Left)
		q.Push(curNode.Right)
	}
	// pop
	//for !q.Empty() {
	//    curNode := q.Top().(*TreeNode)
	//    q.Pop()
	//    curNode.Val = -8848
	//    curNode.Left = nil
	//    curNode.Right = nil
	//}
	root = Trim(root)
	return root
}

func Trim(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == -8848 {
		return nil
	}
	root.Left = Trim(root.Left)
	root.Right = Trim(root.Right)
	return root
}

func (t *TreeNode) PrintTree() {
	q := NewQueue()
	q.Push(t)
	for !q.Empty() {
		curNode := q.Top().(*TreeNode)
		q.Pop()
		fmt.Print(curNode.ToString(), " ")
		if curNode != nil {
			q.Push(curNode.Left)
			q.Push(curNode.Right)
		}
	}
}

func (t *TreeNode) ToString() string {
	if t == nil {
		return "nil"
	}
	return strconv.Itoa(t.Val)
}
