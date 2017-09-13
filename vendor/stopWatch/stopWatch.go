package stopWatch

import (
	"fmt"
	"time"
)

type StopWatch struct{}

var t time.Time

func New() *StopWatch {
	return &StopWatch{}
}
func init() {
	t = time.Now()

}
func (s *StopWatch) Watch() {
	fmt.Println("程序运行时间:	", time.Since(t))
}
