package search

type Search struct{}

// new search
func New() *Search {
	return &Search{}
}

// 二分查找
func (s *Search) BinarySearch(args []int, key int) int {
	lo, hi := 0, len(args)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < args[mid] {
			hi = mid - 1
		} else if key > args[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
