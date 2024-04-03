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

func getPrintMethod(url string) func(string, ...interface{}) {
	var c color.Attribute

	switch {
	case strings.Contains(url, "GET"):
		c = color.BgGreen
	case strings.Contains(url, "POST"):
		c = color.BgYellow
	case strings.Contains(url, "PUT"):
		c = color.BgMagenta
	case strings.Contains(url, "DELETE"):
		c = color.BgRed
	default:
		c = color.BgYellow
	}

	return color.New(c).PrintfFunc()
}

func getPrintCode(statusCode int) func(string, ...interface{}) {
	var colorCode color.Attribute

	switch {
	case statusCode >= 200 && statusCode < 300:
		colorCode = color.FgGreen
	case statusCode >= 300 && statusCode < 400:
		colorCode = color.FgYellow
	default:
		colorCode = color.FgRed
	}

	return color.New(colorCode).PrintfFunc()
}
