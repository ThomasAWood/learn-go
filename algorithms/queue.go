package algorithms

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

type Queue struct {
	list LinkedList
}

func (q *Queue) Peek() int {
	if q.list.head != nil && q.list.tail != nil {
		return q.list.head.value
	} else {
		return -1
	}
}

func (q *Queue) Enqueue(value int) {
	var node = Node{value, nil, nil}
	if q.list.head == nil && q.list.tail == nil {
		q.list.head = &node
		q.list.tail = &node
	} else {
		q.list.tail.next = &node
		q.list.tail = &node
	}
}

func (q *Queue) Dequeue() int {
	if q.list.head != nil {
		var value int = q.list.head.value
		q.list.head = q.list.head.next
		return value
	} else {
		return -1
	}
}
