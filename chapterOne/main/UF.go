package main

import (
	"fmt"
	"readFile"
	"stopWatch"
)

var (
	id    map[int]int
	count int
)

type UFStruct struct{}

// New 创建UFStruct
func New() *UFStruct {
	return &UFStruct{}
}

func main() {
	u := New()
	u.Tiny()
}

// Tiny quick-find算法，时间复杂度为O(n^2)
func (u *UFStruct) Tiny() {
	fmt.Println("UF quick-find begin:		")
	arrayInt, r, fileName := []int{}, readFile.New(), "tinyUF.txt"
	err := r.ReadToInt(fileName, &arrayInt)
	//err := r.ReadToInt("largeUF.txt", &arrayInt) // 具体执行多久不知道，抽了根烟回来还在跑
	if err != nil {
		fmt.Println("error:		", err)
		return
	}
	// 初始化map
	u.UF(&id, arrayInt[0])
	for i := 1; i < len(arrayInt); i += 2 {
		p, q := arrayInt[i], arrayInt[i+1]
		if u.Connected(p, q, fileName) {
			continue
		}
		u.UnionTiny(p, q)
		fmt.Println(p, "---", q)
	}
	fmt.Println(count, "	components")
	stopWatch.New().Watch()
}

// UF 初始化map
func (u *UFStruct) UF(id *map[int]int, n int) {
	count = n
	*id = make(map[int]int)
	for i := 0; i < n; i++ {
		(*id)[i] = i
	}
}

// Connected 检查两个是否相等
func (u *UFStruct) Connected(p, q int, fileName string) bool {
	switch fileName {
	case "tinyUF.txt":
		return u.FindTiny(p) == u.FindTiny(q)
	default:
		return u.FindTiny(p) == u.FindTiny(q)
	}
}

// FindTiny 查找数组值
func (u *UFStruct) FindTiny(p int) int {
	return id[p]
}

// UnionTiny 归并分量操作
func (u *UFStruct) UnionTiny(p, q int) {
	pID, qID := u.FindTiny(p), u.FindTiny(q)

	if pID == qID {
		return
	}
	for k, v := range id {
		if v == pID {
			id[k] = qID
		}
	}
	count--
}
