package main

import (
	. "github.com/qingsong-he/ce"
)

// grouping array
func GroupArray(arraySize int, groupSize int) [][]int {
	var begin, end int

	result := [][]int{}

	if arraySize <= 0 {
		return result
	}

	if groupSize <= 0 {
		groupSize = arraySize
	}

	a := arraySize / groupSize
	d := arraySize % groupSize

	for i := 0; i < a; i++ {
		begin = i * groupSize
		end = (i+1)*groupSize - 1
		result = append(result, []int{begin, end})
	}

	if d != 0 {
		begin = a * groupSize
		end = a*groupSize + d - 1
		result = append(result, []int{begin, end})
	}

	return result
}

func main() {
	Print(GroupArray(10, 2))
	Print(GroupArray(10, 3))
	Print(GroupArray(10, 11))
	Print(GroupArray(0, 11))
	Print(GroupArray(11, 0))
	Print(GroupArray(10, -1))
	Print(GroupArray(-1, 10))
}
