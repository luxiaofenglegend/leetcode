package linkedList

import "leetcode/dataStructure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	length := func(h *ListNode) (len int) {
		len = 0
		for h != nil {
			h = h.Next
			len++
		}
		return
	}(head)
	idx := length - n
	if idx == 0 {
		return head.Next
	} else {
		p := head
		for i := 0; i < idx-1; i++ {
			p = p.Next
		}
		p.Next = p.Next.Next
	}
	return head
}

func Test19() {
	l1 := []int{1, 2, 3, 4, 5}
	list_1 := dataStructure.BuildList(l1)
	list_1.MyPrint()
	res := RemoveNthFromEnd(list_1, 1)
	res.MyPrint()
}
