package main

import (
	. "github.com/qingsong-he/ce"
)

func main() {
	var a []int
	Print(a == nil)      // true
	Print(len(a) == 0)   // true
	Print(a[0:0] == nil) // true
	Print(len(a[0:0]))   // 0
}
