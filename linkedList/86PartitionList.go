package linkedList

import (
	"leetcode/dataStructure"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	s := &ListNode{}
	g := &ListNode{}
	sNode := s
	gNode := g
	for head != nil {
		val := head.Val
		if val < x {
			sNode.Next = &ListNode{val, nil}
			sNode = sNode.Next
		} else {
			gNode.Next = &ListNode{val, nil}
			gNode = gNode.Next
		}
		head = head.Next
	}
	//s.MyPrint()
	//g.MyPrint()
	tmpNode := &ListNode{}
	curNode := tmpNode
	if s.Next != nil {
		curNode.Next = s.Next
		curNode = sNode
		//fmt.Print("ho")
	}
	if g.Next != nil {
		curNode.Next = g.Next
		//fmt.Print("hi")
	}
	//tmpNode.MyPrint()
	return tmpNode.Next
}

func Partition(head *ListNode, x int) *ListNode {
	return partition(head, x)
}

func Test86() {
	l1 := dataStructure.BuildList([]int{1, 4, 3, 2, 5, 2})
	l2 := dataStructure.BuildList([]int{1})
	x := 3
	x2 := 1
	res := partition(l1, x)
	res.MyPrint()
	res2 := partition(l2, x2)
	res2.MyPrint()
}
