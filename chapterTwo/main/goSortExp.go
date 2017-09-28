package main

import (
	"fmt"
	"sort"
)

type Wang struct {
	Age  int
	Name string
}
type WangArray []Wang

func (w WangArray) Len() int {
	return len(w)
}
func (w WangArray) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func (w WangArray) Less(i, j int) bool {
	return w[i].Age < w[j].Age
}

func main() {
	arr := WangArray{{13, "1"}, {15, "3"}, {14, "2"}}
	sort.Sort(arr)
	fmt.Println(arr)
}
func typeOf(v interface{}) interface{} {

	switch t := v.(type) {
	case int:
		fmt.Println(t)
		return "int"
	case float64:
		fmt.Println(t)
		return "float64"
	case string:
		fmt.Println(t)
		return "string"
	default:
		return "unKnow"
	}
}
