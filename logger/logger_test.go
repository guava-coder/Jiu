package logger

import (
	"fmt"
	"testing"
)

func TestPrintlnLatesy(t *testing.T) {
	PrintlnLatesy("test", func() int {
		for i := 0; i < 100; i++ {
			fmt.Print(i)
		}
		return 200
	})
}
