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

// GetIntArray 获取len长度的int数组
// maxInt 取值范围 len int数组长度
func (m *MyRand) GetIntArray(maxInt, len int) []int {
	r, arr := rand.New(rand.NewSource(time.Now().UnixNano())), make([]int, len)
	for i := 0; i < len; i++ {
		arr[i] = r.Intn(maxInt)
	}
	return arr
}

// GetRandomString 随机字符串
// l 字符串长度
func (m *MyRand) GetRandomString(l int) string {
	//str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//bytes := []byte(str)
	result := make([]byte, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; {
		if id := r.Intn(len(letterBytes)); id < len(letterBytes) {
			result[i] = letterBytes[id]
			i++
		}
	}
	return string(result)
}

// GetStringArray 获取随机字符数组
// maxLen 字符串最大长度 len 数组大小
func (m *MyRand) GetStringArray1(maxLen, len int) []string {
	r, arr, str := rand.New(rand.NewSource(time.Now().UnixNano())), []string{}, ""
	for i := 0; i < len; i++ {
		str = m.GetRandomString(r.Intn(maxLen))
		arr = append(arr, str)
	}
	return arr
}

// 用掩码进行替换
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GetStringArray 获取随机字符数组
// maxLen 字符串最大长度 len 数组大小
// 这个方法会比GetStringArray1方法快很多
// 具体表现为生成百万个随机字符串这个方法大概为400ms，上面的方法为12s左右
func (m *MyRand) GetStringArray(maxLen, len int) []string {
	r, arr, str := rand.New(rand.NewSource(time.Now().UnixNano())), []string{}, ""
	for i := 0; i < len; i++ {
		str = m.RandStringBytesMaskImpr(r.Intn(maxLen))
		arr = append(arr, str)
	}
	return arr
}
func (m *MyRand) RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) { // 此处会产生性能差距
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
