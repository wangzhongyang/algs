package stopWatch

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for i := 1; i < 100000; i++ {
		fmt.Print("-")
	}
	new(StopWatch).Watch()
}
