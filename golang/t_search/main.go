package main

import (
	"github.com/qingsong-he/ce"
	"os"
	"sort"
)

func init() {
	ce.Print(os.Args[0])
}

func main() {
	x := 11
	s := []int{3, 6, 8, 11, 45}
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })

	if pos < len(s) && s[pos] == x {
		ce.Print("yes", pos)
	} else {
		ce.Print("no", pos)
	}
}
