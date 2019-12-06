package dataStructure

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.GetLen(), q.Top())
	q.Pop()
	fmt.Println(q.GetLen(), q.Top())
	q.Pop()
}
