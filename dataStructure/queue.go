package dataStructure

type Queue struct {
	arr []interface{}
}

func NewQueue() Queue {
	q := Queue{}
	q.arr = make([]interface{}, 0, 8)
	return q
}

func (q *Queue) GetLen() int {
	return len(q.arr)
}

func (q *Queue) Push(data interface{}) {
	q.arr = append(q.arr, data)
}

func (q *Queue) Pop() {
	length := q.GetLen()
	q.arr = q.arr[1:length]
}

func (q *Queue) Top() interface{} {
	if q.Empty() {
		return nil
	}
	return q.arr[0]
}

func (q *Queue) Empty() bool {
	return q.GetLen() == 0
}
