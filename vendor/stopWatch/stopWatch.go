package stopWatch

import (
	"fmt"
	"time"
)

// StopWatch 创建StopWatch
type StopWatch struct{}

var (
	t    time.Time
	last time.Time
)

// New 新建StopWatch
func New() *StopWatch {
	return &StopWatch{}
}

// init 计时初始化
func init() {
	t = time.Now()
	last = t
	fmt.Println("计时开始：-------")
}

// Watch 距离开始运行的时间
func (s *StopWatch) Watch() {
	last = time.Now()
	fmt.Println("程序距开始运行时间:	", time.Since(t))
}

// LastTime 距离上次计时的时间
func (s *StopWatch) LastTime() {
	fmt.Println("程序距上次运行时间:	", time.Since(t))
	last = time.Now()
}

// Reset 重置计时时间
func (s *StopWatch) Reset() {
	last = time.Now()
}
