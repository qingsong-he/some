package main

import (
	. "github.com/qingsong-he/ce"
	"os"
	"regexp"
	"strings"
)

func Case1() {
	if len(os.Args) == 4 {
		r := regexp.MustCompile(os.Args[2])

		Print(
			r.MatchString(os.Args[3]),
			strings.Join(r.FindAllString(os.Args[3], -1), ","),
		)
	}
}

func main() {
	Case1()
}
