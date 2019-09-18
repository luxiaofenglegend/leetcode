package linkedList

import "leetcode/dataStructure"

type ListNode = dataStructure.ListNode

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
