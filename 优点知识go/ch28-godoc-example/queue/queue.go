package queue

type Queue []interface{}

func (q *Queue) Push(n interface{}) {
	*q = append(*q, n)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
