package main

import (
	. "github.com/qingsong-he/ce"
)

func Limit(n, offset, limit int) [2]int {
	var startIndex, endIndex int

	if offset > n {
		offset = n
	}
	if n < limit {
		limit = n
	}

	startIndex = offset
	if n-startIndex < limit {
		endIndex = startIndex + (n - startIndex)
	} else {
		endIndex = startIndex + limit
	}

	return [2]int{startIndex, endIndex}
}

func main() {
	l := []int{11, 22, 33}
	Print(Limit(0, 2, 6))
	Print(Limit(0, 6, 2))
	Print(Limit(1, 2, 6))
	Print(Limit(1, 6, 2))
	Print(Limit(1, 0, 6))
	Print(Limit(1, 6, 0))
	Print(Limit(10, 0, 2))
	Print(Limit(10, 2, 3))

	interval := Limit(len(l), 0, 2)
	Print(interval)
	Print(l[interval[0]:interval[1]])

	interval = Limit(len(l), 999, 9999)
	Print(interval)
	Print(l[interval[0]:interval[1]])
}
