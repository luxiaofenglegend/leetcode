package linkedList

import "leetcode/dataStructure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*ListNode) *ListNode {
	const INT_MAX = int(^uint(0) >> 1)
	done := false
	//length := len(lists)
	tmpHead := &ListNode{}
	curNode := tmpHead
	cnt := 0
	for !done {
		done = true
		minValue := INT_MAX
		minIndex := -1
		_ = minIndex
		for idx, node := range lists {
			if node != nil {
				if node.Val < minValue {
					minValue = node.Val
					minIndex = idx
					done = false
				}
			}
		}
		if !done {
			curNode.Next = &ListNode{minValue, nil}
			curNode = curNode.Next
			lists[minIndex] = lists[minIndex].Next
			cnt++
		}
	}
	//fmt.Println(cnt)
	return tmpHead.Next
}

func Test23() {
	l1 := dataStructure.BuildList([]int{1, 4, 5})
	l2 := dataStructure.BuildList([]int{1, 3, 4})
	l3 := dataStructure.BuildList([]int{2, 6})
	res := MergeKLists([]*dataStructure.ListNode{l1, l2, l3})
	res.MyPrint()
}
