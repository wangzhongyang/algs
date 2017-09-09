// 作为

package main

import (
	"fmt"
	"queue"
)

func main() {
	//s := stack.New()
	//if s.IsEmpty() {
	//	fmt.Println("true")
	//}
	//for i := 1; i <= 7; i++ {
	//	s.Push(i)
	//}
	//// 访问所有元素
	//s.Foreach()

	q := queue.New()
	for i := 1; i <= 7; i++ {
		q.Enqueue(i)
	}
	for i := 1; i <= 8; i++ {
		if val, ok := q.Dequeue(); ok {
			fmt.Println(val)
		}
	}
}
