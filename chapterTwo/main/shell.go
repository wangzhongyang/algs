// 希尔排序
// 特点: 权衡了子数组的规模和有序性，可用于大数组，对任意排序的数组表现也很好
// 个人感觉是插入排序的改进版，通过构造局部有序数列加快排序速度

package main

import (
	"fmt"
	"myError"
	"readFile"
)

// Shell
type Shell struct{}

// NewShell
func NewShell() *Shell {
	return &Shell{}
}
func main() {
	f, arr, s := readFile.New(), []string{}, NewShell()
	err := f.ReadToStrings("words3.txt", &arr)
	myError.New().Err(err, "读取文件失败")

	s.Sort(&arr)
	fmt.Println("是否为升序：", s.IsSorted(arr))
	fmt.Println(arr)

}

// Sort
func (s *Shell) Sort(args *[]string) {
	n, h := len(*args), 1
	for h < n/3 {
		h = 3*h - 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && s.Less((*args)[j], (*args)[j-h]); j -= h {
				s.Exch(args, j, j-h)
			}
		}
		fmt.Println(args)
		h = h / 3
	}
}

// Less
func (s *Shell) Less(p, q string) bool {
	return !(p > q)
}

// Exch
func (s *Shell) Exch(args *[]string, p, q int) {
	(*args)[p], (*args)[q] = (*args)[q], (*args)[p]
}

// IsSorted
func (s *Shell) IsSorted(args []string) bool {
	for i := 1; i < len(args); i++ {
		if s.Less(args[i], args[i-1]) {
			return false
		}
	}
	return true
}
