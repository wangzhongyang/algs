// 选择排序
// 特点: 1、运行时间和输入无关
// 	    2、数据移动最少
// 复杂度: (n^2)/2次比较和n次交换
package main

import (
	"fmt"
	"myError"
	"readFile"
	"strconv"
)

// Selection
type Selection struct{}

// NewSelection
func NewSelection() *Selection {
	return &Selection{}
}

func main() {
	f, s, arr := readFile.New(), NewSelection(), []string{}
	err := f.ReadToStrings("tiny.txt", &arr)
	myError.New().Err(err, "读取文件失败")

	s.Sort(&arr)
	b := s.IsSorted(arr)
	fmt.Println(b) // 此处输出false，因为有相同的元素
	fmt.Println(arr)

}

// Sort
func (s *Selection) Sort(args *[]string) {
	len := len(*args)
	for i := 0; i < len; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			fmt.Println((*args)[j], (*args)[min])
			if s.Less((*args)[j], (*args)[min]) {
				min = j
			}
		}
		s.Exch(args, i, min)
	}
}

// Less
func (s *Selection) Less(p, q string) bool { // 此处若为数值转为的String，"55"比"7"小
	if a, err := strconv.Atoi(p); err == nil {
		b, _ := strconv.Atoi(q)
		return !(a > b)
	} else {
		return !(p > q)
	}
}

// Exch
func (s *Selection) Exch(args *[]string, p, q int) {
	(*args)[p], (*args)[q] = (*args)[q], (*args)[p]
}

// IsSorted
func (s *Selection) IsSorted(args []string) bool {
	for i := 1; i < len(args); i++ {
		if s.Less(args[i], args[i-1]) {
			return false
		}
	}
	return true
}
