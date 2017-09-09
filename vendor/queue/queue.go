package queue

// 队列
type Queue struct {
	first *Element
	last  *Element
	size  int
}
type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

// new queue
func New() *Queue {
	return &Queue{}
}

// get queue size
func (q *Queue) Len() int {
	return q.size
}

// Return the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
func (q *Queue) Enqueue(value interface{}) {
	e := &Element{value: value, next: nil}
	if q.IsEmpty() {
		q.first, q.last = e, e
	} else {
		old := q.last
		q.last = e
		old.next = q.last
		//q.last, old.next = e, q.last
	}
	q.size++
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	value := q.first.value
	q.first = q.first.next
	q.size--
	return value, true
}

// Foreach queue
func (q *Queue) Foreach() {
	for q.Len() > 0 {
		q.first = q.first.next
		q.size--
	}
}
