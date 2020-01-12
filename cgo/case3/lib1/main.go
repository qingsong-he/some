package main

import "C"

// go install -buildmode=shared -linkshared std

//export Sum
func Sum(x, y int) int {
	return x + y
}

func main() {
}
