package queuemap

type queueNode struct {
	Prev     *queueNode
	Next     *queueNode
	KeyValue KeyValue
}

type queue struct {
	Head *queueNode
	Tail *queueNode
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) Push(qn *queueNode) {
	if q.Head == nil {
		q.Head = qn
		q.Tail = qn
		qn.Prev = nil
		qn.Next = nil
	} else {
		q.Tail.Next = qn
		qn.Prev = q.Tail
		q.Tail = qn
	}
}

func (q *queue) Remove(qn *queueNode) {
	if qn.Prev == nil {
		q.Head = qn.Next
	} else {
		qn.Prev.Next = qn.Next
	}
	if qn.Next == nil {
		q.Tail = qn.Prev
	} else {
		qn.Next.Prev = qn.Prev
	}
}
