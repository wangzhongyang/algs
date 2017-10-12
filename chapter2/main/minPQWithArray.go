package main

import (
	"fmt"
	"myRand"
	"strconv"
)

// 数组实现的二叉堆 及优先队列
type MinPQWithArray struct {
	Array []string
	Len   int
}

func main() {
	m, arrInt := NewMinPQWithArray(), []int{}
	m.Min(arrInt)
}

// NewMaxPQWithArray
func NewMinPQWithArray() *MinPQWithArray {
	return &MinPQWithArray{}
}

// min 优先队列，由小到大
func (m *MinPQWithArray) Min(arrInt []int) {
	m, maxType := NewMinPQWithArray(), MinPQWithArray{
		Array: []string{""},
		Len:   0,
	}
	if len(arrInt) == 0 {
		arrInt = myRand.New().GetIntArray(1000, 20)
	}
	fmt.Println(arrInt)
	for _, v := range arrInt {
		m.Insert(&maxType, strconv.Itoa(v))
		fmt.Println(maxType)
	}

	for maxType.Len > 0 {
		fmt.Print("max string:		", m.DelMax(&maxType), "	")
		fmt.Println(maxType)
	}
}

// IsEmpty
func (m *MinPQWithArray) IsEmpty() bool {
	return m.Len == 1
}

// Size
func (m *MinPQWithArray) Size() int {
	return m.Len
}

// Insert
func (m *MinPQWithArray) Insert(args *MinPQWithArray, s string) {
	args.Len += 1
	args.Array = append(args.Array, s)
	m.Swim(args)
}

// DelMax
func (m *MinPQWithArray) DelMax(args *MinPQWithArray) string {
	maxStr := args.Array[1]                   // 最大元素
	m.Exch(&args.Array, 1, len(args.Array)-1) // 交换首尾
	args.Array = args.Array[:args.Len]        // 裁剪最后一个元素
	args.Len--                                // 长度减一
	m.Sink(&args.Array, 1)                    // 恢复堆的有序性
	return maxStr
}

//////////////////
// 以下为辅助函数 //
//////////////////
// Sink 由上至下的堆有序化（下沉）
func (m *MinPQWithArray) Sink(args *[]string, k int) {
	arr, l := *args, len(*args)
	for k*2 < l {
		j := k * 2
		if j+1 < l && m.Less(arr[j], arr[j+1]) {
			j++
		}
		if !(m.Less(arr[k], arr[j])) {
			break
		}
		m.Exch(args, k, j)
		k = j
	}
}

// Swim 由下至上的堆有序化（上浮）
func (m *MinPQWithArray) Swim(args *MinPQWithArray) {
	arr, k := args.Array, args.Len
	for k > 1 && m.Less(arr[k/2], arr[k]) {
		m.Exch(&arr, k/2, k)
		k = k / 2
	}
}

// Less
func (m *MinPQWithArray) Less(p, q string) bool { // 此处若为数值转为的String，"55"比"7"小
	if a, err := strconv.Atoi(p); err == nil {
		b, _ := strconv.Atoi(q)
		return !(a < b)
	} else {
		return !(p < q)
	}
}

// Exch
func (m *MinPQWithArray) Exch(args *[]string, p, q int) {
	(*args)[p], (*args)[q] = (*args)[q], (*args)[p]
}
