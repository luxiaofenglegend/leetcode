package linkedList

import "leetcode/dataStructure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs(head *ListNode) *ListNode {
	tmpHead := &ListNode{}
	curNode := tmpHead
	for head != nil && head.Next != nil {
		firstValue := head.Val
		secondValue := head.Next.Val
		curNode.Next = &ListNode{secondValue, nil}
		curNode = curNode.Next
		curNode.Next = &ListNode{firstValue, nil}
		curNode = curNode.Next
		head = head.Next.Next
	}
	if head != nil {
		curNode.Next = &ListNode{head.Val, nil}
	}
	return tmpHead.Next
}
func Test24() {
	l1 := dataStructure.BuildList([]int{1, 2, 3, 4, 5})
	res := SwapPairs(l1)
	res.MyPrint()
}
