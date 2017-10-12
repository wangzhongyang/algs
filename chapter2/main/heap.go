package main

import (
	"fmt"
	"strconv"
)

type Heap struct{}

func NewHeap() *Heap {
	return &Heap{}
}
func main() {
	h, arr := NewHeap(), []string{"", "S", "O", "R", "T", "E", "X", "A", "E", "P", "L", "E"}

	h.Sort(&arr)
	fmt.Println(arr)
}
func (h *Heap) Sort(args *[]string) {
	l := len(*args) - 1
	// 堆的有序化
	for k := l / 2; k > 1; k-- {
		h.Sink(args, k, l)
	}
	fmt.Println(args)
	// 排序
	for l > 1 {
		h.Exch(args, 1, l)
		l--
		h.Sink(args, 1, l)
	}
}

// Sink 下沉
func (h *Heap) Sink(args *[]string, k, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && j+1 < n && h.Less((*args)[j], (*args)[j+1]) {
			j++
		}
		if !h.Less((*args)[k], (*args)[j]) {
			break
		}
		h.Exch(args, k, j)
		k = j
	}
}

// Less
func (h *Heap) Less(p, q string) bool { // 此处若为数值转为的String，"55"比"7"小
	if a, err := strconv.Atoi(p); err == nil {
		b, _ := strconv.Atoi(q)
		return !(a < b)
	} else {
		return !(p < q)
	}
}

// Exch
func (h *Heap) Exch(args *[]string, p, q int) {
	(*args)[p], (*args)[q] = (*args)[q], (*args)[p]
}
