package tools

import (
	"fmt"
	"strings"
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

	printMethod := getPrintMethod(url)
	printMethod("%s ", url)

	printCode := getPrintCode(code)
	printCode("| %d |---> %s", code, time.Since(start))

	fmt.Println()
}

func getPrintMethod(url string) (f func(string, ...interface{})) {
	f = color.New(color.BgYellow).PrintfFunc()

	if strings.Contains(url, "GET") {
		f = color.New(color.BgGreen).PrintfFunc()
	}
	return
}

func getPrintCode(code int) (f func(string, ...interface{})) {
	f = color.New(color.FgRed).PrintfFunc()

	if code >= 200 && code < 300 {
		f = color.New(color.FgGreen).PrintfFunc()
	}
	return
}
