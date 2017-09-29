// 归并排序
// 归并排序是一种渐进最优的基于比较排序的算法
// 空间复杂度不是最优
package main

import (
	"fmt"
	"myError"
	"readFile"
	"strconv"
)

// Merge
type Merge struct{}

// NewMerge
func NewMerge() *Merge {
	return &Merge{}
}
func main() {
	m, fileName := NewMerge(), "words3.txt"
	//m.UpToDown(fileName)
	m.DownToUp(fileName)
}

// DownToUp 自底向上的归并排序
// 特点:
// 		比较适合用链表组织的数据
func (m *Merge) DownToUp(fileName string) {
	f, arr := readFile.New(), []string{}
	err := f.ReadToStrings(fileName, &arr)
	myError.New().Err(err, "读取文件失败")

	fmt.Println(arr)
	m.SortBU(&arr)
	fmt.Println("是否为升序：		", m.IsSorted(arr))
	fmt.Println(arr)
}

// UpToDown 自顶向下的归并排序排序
// 特点:1、所需时间和NlgN成正比
// 	   2、辅助数组锁使用的额外空间和N的大小成正比
// 比较次数为: NlgN这里就不推导了，看书吧
func (m *Merge) UpToDown(fileName string) {
	f, arr, aux := readFile.New(), []string{}, []string{}
	err := f.ReadToStrings(fileName, &arr)
	myError.New().Err(err, "读取文件失败")
	aux = make([]string, len(arr))

	fmt.Println(arr)
	m.SortUD(&arr, &aux, 0, len(arr)-1)
	fmt.Println("是否为升序：		", m.IsSorted(arr))
	fmt.Println(arr)
}

// SortUD 自顶向下的归并排序
func (m *Merge) SortUD(args, aux *[]string, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	m.SortUD(args, aux, lo, mid)
	m.SortUD(args, aux, mid+1, hi)
	m.Merge(args, aux, lo, mid, hi)
}

// SortBU 自底向上的归并排序
func (m *Merge) SortBU(args *[]string) {
	n, aux := len(*args), make([]string, len(*args))
	for sz := 1; sz < n; sz *= 2 {
		for lo := 0; lo < n-sz; lo += sz + sz {
			mid, hi := lo+sz-1, m.Min(lo+sz+sz-1, n-1)
			m.Merge(args, &aux, lo, mid, hi)
		}
	}
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
func (m *Merge) Less(p, q string) bool { // 此处若为数值转为的String，"55"比"7"小
	if a, err := strconv.Atoi(p); err == nil {
		b, _ := strconv.Atoi(q)
		return !(a > b)
	} else {
		return !(p > q)
	}
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

// Min
func (m *Merge) Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
