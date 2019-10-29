package linkedList

import "leetcode/dataStructure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmpHead := &ListNode{
		Val:  0,
		Next: head,
	}
	last_val := head.Val
	last := head
	cur := head.Next
	for cur != nil {
		if cur.Val == last_val {
			last.Next = cur.Next
		} else {
			last_val = cur.Val
			last = cur
		}
		cur = cur.Next
	}
	return tmpHead.Next
}

func Test83() {
	l1 := []int{1, 1, 2, 3, 3}
	list_1 := dataStructure.BuildList(l1)
	list_1.MyPrint()

	res := DeleteDuplicates(list_1)
	res.MyPrint()
}
