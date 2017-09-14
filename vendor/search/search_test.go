package search

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	s := new(Search)
	args := []int{1, 2, 3, 4, 5, 6, 23, 67, 76, 24, 12}
	sort.Ints(args)
	fmt.Println(args)
	key := s.BinarySearch(args, 76)
	fmt.Println("key is:		", key)
	if key != -1 {
		fmt.Println("value is:	", args[key])
	}
	key = s.BinarySearch(args, 760)
	fmt.Println("key is:		", key)
	if key != -1 {
		fmt.Println("value is:	", args[key])
	}
}
