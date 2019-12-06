package dataStructure

import "fmt"

type myPrint interface {
	myPrint()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) MyPrint() {
	for ln != nil {
		fmt.Print(ln.Val, " ")
		ln = ln.Next
	}
	fmt.Println()
}

func BuildList(a []int) *ListNode {
	res := new(ListNode)
	cur := res
	for _, val := range a {
		cur.Next = new(ListNode)
		cur = cur.Next
		cur.Val = val
	}
	return res.Next
}
