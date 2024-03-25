package tools

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func PrintlnLatency(url string, request func() int) {
	start := time.Now()
	fmt.Println(start)

	code := request()

	printGet := color.New(color.FgGreen).PrintfFunc()
	printGet("%s | %d |---> %s", url, code, time.Since(start))
	fmt.Println()
}
