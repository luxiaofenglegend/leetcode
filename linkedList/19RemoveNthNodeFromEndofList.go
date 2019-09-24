package linkedList

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
