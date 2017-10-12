package main

import "fmt"

// 查找示例代码结构
type Example struct{}

var args map[string]string

func main() {
	fmt.Println("8888")
}

// NewExample get new example
func NewExample() *Example {
	return &Example{}
}

//Put 将键值对存入
func (e *Example) Put(key, value string) {
	args[key] = value
}

// Get 获取对应的值，不存在返回空
func (e *Example) Get(key string) string {
	result, ok := args[key]
	if ok {
		return result
	} else {
		return "NaN"
	}
}

//Delete 删除key
func (e *Example) Delete(key string) {
	delete(args, key)
}

//Contains 键是否存在
func (e *Example) Contains(key string) bool {
	return e.Get(key) != "NaN"
}

//IsEmpty 是否为空
func (e *Example) IsEmpty() bool {
	return e.Size() == 0
}

//Size 键值对的数量
func (e *Example) Size() int {
	return 0
}

//Min 最小的键
func (e *Example) Min() string {
	return ""
}

//Max 最大的键
func (e *Example) Max() string {
	return ""
}

//Floor 小于等于key的最大键
func (e *Example) Floor(key string) string {
	return ""
}

//Ceiling 大于等于key的最小键
func (e *Example) Ceiling(key string) string {
	return ""
}

//Rank 小于key的数量
func (e *Example) Rank(key string) int {
	return 0
}

//Select 排名为n的键
func (e *Example) Select(n int) string {
	return ""
}

//DeleteMin 删除最小的键
func (e *Example) DeleteMin() {

}

//DeleteMax 删除最大的键
func (e *Example) DeleteMax() {

}

//LHSize [lo...hi]之间键的数量
func (e *Example) LHSize() int {
	return 0
}

//LHKeys [lo...hi]之间的所有键，已排序
func (e *Example) LHKeys() []string {
	return []string{}
}

// Keys 所有键的集合，已排序
func (e *Example) Keys() []string {
	return []string{}
}
