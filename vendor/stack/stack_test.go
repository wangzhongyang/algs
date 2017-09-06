package stack

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	stack := new(Stack)

	stack.Push("Things")
	stack.Push("and")
	stack.Push("Stuff")

	fmt.Printf("%s ", stack.Pop().(string))
	stack.Push(6.4)
	for stack.Len() > 0 {
		// We have to do a type assertion because we get back a variable of type
		// interface{} while the underlying type is a string.
		fmt.Printf("%s ", stack.Pop().(string))
	}
}
