package main

import (
	. "github.com/qingsong-he/ce"
	"unsafe"
)

func Case1() {
	var f float64
	var i int64

	// any pointer type cast to unsafe.Pointer
	var p unsafe.Pointer
	p = unsafe.Pointer(&f)
	Print(p)

	p = unsafe.Pointer(&i)
	Print(p)

	// unsafe.Pointer cast to any pointer type
	Print((*int)(p))

	// unsafe.Pointer cast to uintptr
	var k uintptr
	k = uintptr(p)
	Print(k)

	// uinptr cast to unsafe.Pointer
	p = unsafe.Pointer(k)
	Print(p)
}

func main() {
	Case1()
}
