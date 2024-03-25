package tools

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// PrintlnLatency prints the latency of a request to the specified URL.
//
// Parameters:
// - url: the URL to send the request to.
// - request: a function that performs the request and returns the status code.
func PrintlnLatency(url string, request func() int) {
	start := time.Now()
	fmt.Println(start)

	code := request()

	printGet := color.New(color.FgGreen).PrintfFunc()
	printGet("%s | %d |---> %s", url, code, time.Since(start))
	fmt.Println()
}
