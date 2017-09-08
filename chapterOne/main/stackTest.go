package main

import (
	"stack"
)

func main() {
	s := stack.New()
	for i := 1; i <= 7; i++ {
		s.Push(i)
	}
	// 访问所有元素
	s.Foreach()
	//for s.Len() > 0 {
	//	fmt.Println(s.Pop())
	//}
}
