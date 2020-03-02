package main

import (
	"github.com/qingsong-he/ce"
	"os"
)

func init() {
	ce.Print(os.Args[0])
}

// get [begin, end) of total
func GetBetweenIndex(total, offset, limit int) []int {
	result := []int{}

	if total <= 0 {
		return result
	}

	if offset < 0 || limit <= 0 {
		return result
	}

	// too strict?
	if offset > total-1 {
		return result
	}

	// all := [2]int{0, total}
	// sub := [2]int{offset, offset + limit}

	var startIndex, endIndex int

	// get startIndex
	if 0 <= offset && offset <= total-1 {
		startIndex = offset
	} else {
		startIndex = total - 1
	}

	// get endIndex
	if offset+limit <= total {
		endIndex = offset + limit
	} else {
		endIndex = total
	}

	return []int{startIndex, endIndex}
}

func main() {
	ce.Print(GetBetweenIndex(0, 0, 1))   // []int{}
	ce.Print(GetBetweenIndex(1, 0, 0))   // []int{}
	ce.Print(GetBetweenIndex(1, 0, 1))   // []int{0, 1}
	ce.Print(GetBetweenIndex(1, 1, 1))   // []int{}
	ce.Print(GetBetweenIndex(9, 8, 100)) // []int{8, 9}
	ce.Print(GetBetweenIndex(9, 0, 100)) // []int{0, 9}
}
