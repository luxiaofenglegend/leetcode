package linkedList

import "leetcode/dataStructure"

//type ListNode = dataStructure.ListNode

//Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		res := new(ListNode)
		cur := res
		for l1 != nil && l2 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
			if l1.Val <= l2.Val {
				cur.Val = l1.Val
				l1 = l1.Next
			} else {
				cur.Val = l2.Val
				l2 = l2.Next
			}
		}
		if l1 != nil {
			cur.Next = l1
		} else if l2 != nil {
			cur.Next = l2
		}
		return res.Next
	}
}

func Test21() {
	l1 := []int{1, 2, 4}
	list_1 := dataStructure.BuildList(l1)
	list_1.MyPrint()
	l2 := []int{1, 3, 4}
	list_2 := dataStructure.BuildList(l2)
	list_2.MyPrint()
	res := MergeTwoLists(list_1, list_2)
	res.MyPrint()
}
