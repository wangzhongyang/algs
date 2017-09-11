package queue

import "fmt"

// 队列 有官方实现库
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
	first, len := q.first, q.size
	for len > 0 {
		fmt.Println("遍历不改变结构：", first.value)
		first = first.next
		len--
	}
}
