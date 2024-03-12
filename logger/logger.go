package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Println(url string, request func()) {
	start := time.Now()
	fmt.Println(start)

	request()

	printGet := color.New(color.FgGreen).PrintfFunc()
	printGet("%s |---> %s", url, time.Since(start))
	fmt.Println()
}
