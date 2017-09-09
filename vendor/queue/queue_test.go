package queue

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	q := new(Queue)
	for i := 1; i <= 7; i++ {
		q.Enqueue(i)
	}
	for i := 1; i <= 8; i++ {
		if val, ok := q.Dequeue(); ok {
			fmt.Println(val)
		}
	}
}
