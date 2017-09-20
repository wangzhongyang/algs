package main

import (
	"fmt"
	"readFile"
	"stopWatch"
)

var (
	id       map[int]int
	sz       map[int]int
	arrayInt []int
	count    int
)

type UFStruct struct{}

// New 创建UFStruct
func New() *UFStruct {
	return &UFStruct{}
}

// main union-find算法
func main() {
	u, r, fileName := New(), readFile.New(), "mediumUF.txt"
	err := r.ReadToInt(fileName, &arrayInt)
	//err := r.ReadToInt("largeUF.txt", &arrayInt) // 具体执行多久不知道，抽了根烟回来还在跑
	if err != nil {
		fmt.Println("error:		", err)
		return
	}
	stopWatch.New().Watch()
	u.Tiny(fileName)
	u.Union(fileName)
	u.Last(fileName)
}

// Tiny quick-find算法，时间复杂度为O(n^2)
func (u *UFStruct) Tiny(fileName string) {
	fmt.Println("UF quick-find begin:		")

	// 初始化map
	u.UF(&id, arrayInt[0])
	for i := 1; i < len(arrayInt); i += 2 {
		p, q := arrayInt[i], arrayInt[i+1]
		if u.Connected(p, q, fileName) {
			continue
		}
		u.UnionTiny(p, q)
		//fmt.Println(p, "---", q)
	}
	fmt.Println(count, "	components")
	stopWatch.New().LastTime()
}

// Union 此处未剥离代码主要是为了比较执行时间，仅从tinyUF.txt和mediumUF.txt比较，并没有更快
func (u *UFStruct) Union(fileName string) {
	fmt.Println("UF quick-union begin:		")

	// 初始化map
	u.UF(&id, arrayInt[0])
	for i := 1; i < len(arrayInt); i += 2 {
		p, q := arrayInt[i], arrayInt[i+1]
		if u.Connected(p, q, fileName) {
			continue
		}
		u.UnionUnion(p, q)
		//fmt.Println(p, "---", q)
	}
	fmt.Println(count, "	components")
	stopWatch.New().LastTime()
}

// Last 加权 quick-find 算法
func (u *UFStruct) Last(fileName string) {
	fmt.Println("UF Last quick-union begin:		")

	// 初始化map
	u.LastUF(&id, &sz, arrayInt[0])
	for i := 1; i < len(arrayInt); i += 2 {
		p, q := arrayInt[i], arrayInt[i+1]
		if u.Connected(p, q, fileName) {
			continue
		}
		u.UnionUnion(p, q)
		//fmt.Println(p, "---", q)
	}
	fmt.Println(count, "	components")
	stopWatch.New().LastTime()
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

// FindUnion
func (u *UFStruct) FindUnion(p int) int {
	for p != id[p] {
		p = id[p]
	}
	return p
}

// UnionUnion
func (u *UFStruct) UnionUnion(p, q int) {
	// 将p和q的根节点统一
	pRoot, qRoot := u.FindUnion(p), u.FindUnion(q)
	if pRoot == qRoot {
		return
	}
	id[pRoot] = qRoot
	count--
}

// LastUF 初始化map
func (u *UFStruct) LastUF(id, sz *map[int]int, n int) {
	count = n
	*id, *sz = make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		(*id)[i], (*sz)[i] = i, 1
	}
}

// LastFind
func (u *UFStruct) LastFind(p int) int {
	for p != id[p] {
		p = id[p]
	}
	return p
}

// LastUnion
func (u *UFStruct) LastUnion(p, q int) {
	i, j := u.LastFind(p), u.LastFind(q)
	if i == j {
		return
	}
	if sz[i] < sz[j] {
		id[i], sz[j] = j, sz[j]+sz[i]
	} else {
		id[j], sz[i] = i, sz[i]+sz[j]
	}
	count--
}
