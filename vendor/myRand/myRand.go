package myRand

import (
	"math/rand"
	"time"
)

// MyRand
type MyRand struct{}

// New
func New() *MyRand {
	return &MyRand{}
}

// ShuffleInt 对int数组洗牌
func (m *MyRand) ShuffleInt(args *[]int) []int {
	arr, r := make([]int, len(*args)), rand.New(rand.NewSource(time.Now().Unix()))
	for k, i := range r.Perm(len(*args)) {
		arr[k] = (*args)[i]
	}
	return arr
}

// ShuffleString 对String数组洗牌
func (m *MyRand) ShuffleString(args []string) []string {
	arr, r := make([]string, len(args)), rand.New(rand.NewSource(time.Now().Unix()))
	for k, i := range r.Perm(len(args)) {
		arr[k] = (args)[i]
	}
	return arr
}
