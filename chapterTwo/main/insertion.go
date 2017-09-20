// 插入排序
// 特点: 运行时间取决于输入元素的初始顺序
// 复杂度:
// 		平均 (n^2)/4 比较、交换
// 		最好 n-1次比较 0次交换 		（已经是顺序的数组）
// 		最差 (n^2)/2 比较、交换 	（对角线前数据全要移动）
// 适用情况:
// 		数组中倒置数量小于数组大小的某个倍数很有效
// 		若倒置数量很少时，插入排序是很快的
// 倒置数量：若有 E X A M P L E 中有11个倒置对:E-A、X-A、X-M、X-P、X-L、X-E、M-L、
// 			M-E、P-L、P-E、L-E
package main

import (
	"fmt"
	"myError"
	"readFile"
)

// Insertion
type Insertion struct{}

// NewInsertion
func NewInsertion() *Insertion {
	return &Insertion{}
}

func main() {
	f, arr, i := readFile.New(), []string{}, NewInsertion()
	err := f.ReadToStrings("tiny.txt", &arr)
	myError.New().Err(err, "读取文件失败")

	i.Sort(&arr)
	i.IsSorted(arr)
	fmt.Println(arr)

}

// Sort
func (i *Insertion) Sort(args *[]string) {
	n := len(*args)
	for p := 1; p < n; p++ {
		for j := p; j > 0 && i.Less((*args)[j], (*args)[j-1]); j-- {
			i.Exch(args, j, j-1)
		}
	}
}

// Less
func (i *Insertion) Less(p, q string) bool {
	return !(p > q)
}

// Exch
func (i *Insertion) Exch(args *[]string, p, q int) {
	(*args)[p], (*args)[q] = (*args)[q], (*args)[p]
}

// IsSorted
func (i *Insertion) IsSorted(args []string) bool {
	for p := 1; p < len(args); p++ {
		if i.Less(args[p], args[p+1]) {
			return false
		}
	}
	return true
}
