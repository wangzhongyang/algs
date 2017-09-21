// 自顶向下的归并排序排序
// 特点:1、所需时间和NlgN成正比
// 	   2、辅助数组锁使用的额外空间和N的大小成正比
// 比较次数为: NlgN这里就不推导了，看书吧
package main

import (
	"fmt"
	"myError"
	"readFile"
)

type Merge struct{}

func NewMerge() *Merge {
	return &Merge{}
}
func main() {
	f, arr, aux, m := readFile.New(), []string{}, []string{}, NewMerge()
	err := f.ReadToStrings("tiny.txt", &arr)
	myError.New().Err(err, "读取文件失败")
	aux = make([]string, len(arr))

	m.Sort(&arr, &aux, 0, len(arr)-1)
	fmt.Println("是否为升序：		", m.IsSorted(arr))
	fmt.Println(arr)

}

// Sort
func (m *Merge) Sort(args, aux *[]string, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	m.Sort(args, aux, lo, mid)
	m.Sort(args, aux, mid+1, hi)
	m.Merge(args, aux, lo, mid, hi)
}

// Merge 原地归并的抽象方法
func (m *Merge) Merge(args, aux *[]string, lo, mid, hi int) {
	i, j := lo, mid+1
	// ----------------------------------------------
	// TODO 此处不能这样写，否则*aux会成为 (*args)[lo:hi] 引用，即 (*args)[lo:hi]改变*aux也会改变,采用循环赋值
	//*aux = (*args)[lo:hi]
	// ----------------------------------------------
	for k := i; k <= hi; k++ {
		(*aux)[k] = (*args)[k]
	}
	for k := lo; k <= hi; k++ {
		if i > mid {
			(*args)[k] = (*aux)[j]
			j++
		} else if j > hi {
			(*args)[k] = (*aux)[i]
			i++
		} else if m.Less((*aux)[j], (*aux)[i]) {
			(*args)[k] = (*aux)[j]
			j++
		} else {
			(*args)[k] = (*aux)[i]
			i++
		}
		fmt.Println(args)
	}
}

// Less
func (m *Merge) Less(p, q string) bool {
	return !(p > q)
}

// Exch
func (m *Merge) Exch(args *[]string, p, q int) {
	(*args)[p], (*args)[q] = (*args)[q], (*args)[p]
}

// IsSorted
func (m *Merge) IsSorted(args []string) bool {
	for i := 1; i < len(args); i++ {
		if m.Less(args[i], args[i-1]) {
			return false
		}
	}
	return true
}
