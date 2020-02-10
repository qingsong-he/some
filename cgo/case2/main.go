package main

//#cgo linux CFLAGS: -I./so
//#cgo linux LDFLAGS: -L./so -lso -Wl,-rpath,$ORIGIN/so
//#include "so.h"
import "C"

import "fmt"

func main() {
	a := C.int(1)
	b := C.int(2)
	value := C.sum(a, b)
	fmt.Println(value)
}
