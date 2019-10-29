package linkedList

import "leetcode/dataStructure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	fastPointer := head
	slowPointer := head
	for {
		if fastPointer != nil {
			fastPointer = fastPointer.Next
		}
		if fastPointer != nil {
			fastPointer = fastPointer.Next
		}
		if slowPointer != nil {
			slowPointer = slowPointer.Next
		}

		if (fastPointer == slowPointer) || (fastPointer == nil || slowPointer == nil) {
			break
		}
	}
	// debug
	//if slowPointer != nil {
	//   println(slowPointer.Val)
	//}
	//if fastPointer != nil {
	//   println(fastPointer.Val)
	//}
	if fastPointer == nil || slowPointer == nil {
		return false
	} else {
		return true
	}
	//return false
}

func Test141() {
	l1 := dataStructure.BuildList([]int{1, 4, 3, 2})
	l2 := dataStructure.BuildList([]int{1})
	head := &ListNode{1, &ListNode{2, nil}}
	head.Next.Next = head
	println(hasCycle(l1), hasCycle(l2))
	println(hasCycle(head))
}
