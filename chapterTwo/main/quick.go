// 快速排序
// 特点:1、内循环简洁
// 	   2、比较次数少
// 开始前打乱数组元素顺序为了排除输入对排序的影响
// 优化:1、在小数组时改用插入排序2、三取样切分（即取个中位数）
package main

import (
	"fmt"
	"myError"
	"myRand"
	"readFile"
)

// Quick
type Quick struct{}

// NewQuick
func NewQuick() *Quick {
	return &Quick{}
}

func main() {
	f, arr, qk := readFile.New(), []string{}, NewQuick()
	err := f.ReadToStrings("words3.txt", &arr)
	myError.New().Err(err, "读取文件失败")

	qk.Sort(&arr)
	fmt.Println("是否为升序：", qk.IsSorted(arr))
	fmt.Println(arr)

}

// Sort
func (qk *Quick) Sort(args *[]string) {
	*args = myRand.New().ShuffleString(*args)
	qk.sort(args, 0, len(*args)-1)
}

// sort
func (qk *Quick) sort(args *[]string, lo, hi int) {
	fmt.Println(args)
	if hi <= lo {
		return
	}
	j := qk.Partition(args, lo, hi)
	qk.sort(args, lo, j-1)
	qk.sort(args, j+1, hi)
}

// Partition 快速排序的切分
func (qk *Quick) Partition(args *[]string, lo, hi int) int {
	i, j, v := lo, hi+1, (*args)[lo]
	for true {
		for qk.Less((*args)[qk.Hincr(&i, 1)], v) {
			if i == hi {
				break
			}
		}
		for qk.Less(v, (*args)[qk.Hincr(&j, -1)]) {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		qk.Exch(args, i, j)
	}
	qk.Exch(args, lo, j)
	return j
}

// Less
func (qk *Quick) Less(p, q string) bool {
	return !(p > q)
}

// Exch
func (qk *Quick) Exch(args *[]string, p, q int) {
	(*args)[p], (*args)[q] = (*args)[q], (*args)[p]
}

// IsSorted
func (qk *Quick) IsSorted(args []string) bool {
	for i := 1; i < len(args); i++ {
		if qk.Less(args[i], args[i-1]) {
			return false
		}
	}
	return true
}

// Hincr 增加值
func (qk *Quick) Hincr(args *int, hir int) int {
	*args += hir
	return *args
}
