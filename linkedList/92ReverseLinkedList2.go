package linkedList

import "leetcode/dataStructure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	tmpHead := &ListNode{0, head}
	tmp := head
	last := tmpHead
	for i := 0; i < m-1; i++ {
		last = tmp
		tmp = tmp.Next
	}
	reverseHead := tmp
	for j := m - 1; j < n; j++ {
		tmp = tmp.Next
	}
	//print("tmp: ")
	//tmp.MyPrint()
	len := n - m + 1
	subHead, subTail := reverseListWithHeadAndTail(reverseHead, len)

	last.Next = subHead
	subTail.Next = tmp
	return tmpHead.Next
}

func reverseListWithHeadAndTail(head *ListNode, len int) (h *ListNode, tail *ListNode) {
	if len <= 0 {
		return nil, nil
	}
	tmpHead := &ListNode{0, nil}
	iter := head
	tail = head
	for i := 0; i < len; i++ {
		next := iter.Next
		iter.Next = tmpHead.Next
		tmpHead.Next = iter
		iter = next
	}
	tail.Next = nil
	return tmpHead.Next, tail
}

func Test92() {
	l1 := dataStructure.BuildList([]int{1, 2, 3, 4, 5})
	l1.MyPrint()
	//res, tail := reverseListWithHeadAndTail(l1.Next, 3)
	//tail.Next = nil
	res := reverseBetween(l1, 1, 4)
	res.MyPrint()
}
