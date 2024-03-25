package tools

import (
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	t.Run("test print line latency", func(t *testing.T) {
		PrintlnLatency("test", func() int {
			for i := 0; i < 100; i++ {
				fmt.Print(i)
			}
			return 200
		})

	})
}
