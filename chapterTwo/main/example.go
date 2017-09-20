// 选择排序
// 特点:1、运行时间和输入无关
// 	   2、数据移动最少
package main

import (
	"fmt"
	"myError"
	"readFile"
)

func main() {
	f, arr := readFile.New(), []string{}
	err := f.ReadToStrings("tiny.txt", &arr)
	myError.New().Err(err, "读取文件失败")

	Sort(&arr)
	IsSorted(arr)
	fmt.Println(arr)

}

// Sort
func Sort(args *[]string) {

}

// Less
func Less(p, q string) bool {
	return !(p > q)
}

// Exch
func Exch(args *[]string, p, q int) {
	(*args)[p], (*args)[q] = (*args)[q], (*args)[p]
}

// IsSorted
func IsSorted(args []string) bool {
	for i := 1; i < len(args); i++ {
		if Less(args[i], args[i+1]) {
			return false
		}
	}
	return true
}