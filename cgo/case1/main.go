package main

//#include "so.h"
import "C"

import "fmt"

func main() {
	a := C.int(1)
	b := C.int(2)
	value := C.sum(a, b)
	fmt.Println(value)
}
