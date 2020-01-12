package main

import (
	. "github.com/qingsong-he/ce"
	"os"
	"reflect"
	"unsafe"
)

func Case1() {
	var i1 interface{}

	Print(reflect.TypeOf(i1) == nil)
	Print(reflect.ValueOf(i1).IsValid())
}

func Case2() {
	var i1 interface{}
	i1 = (*int)(nil)

	t1 := reflect.TypeOf(i1)
	Print(t1.Kind() == reflect.Ptr)

	t2 := t1.Elem()
	v1 := reflect.New(t2).Interface().(*int)
	Print(v1, *v1)
}

func Case3() {
	var i1 interface{}

	i1 = false
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)

	i1 = `hello1`
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)

	i1 = 1
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)

	i1 = uint(1)
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)

	i1 = uintptr(1)
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)

	i1 = 1.1
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)

	i1 = 1 + 1i
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)

	i1 = [3]int{1, 2, 3}
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)
	type v1 [3]int
	var v11 v1
	Print(`kind is:`, reflect.TypeOf(v11).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(v11).Name() == ``)

	i1 = make(chan int)
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)
	type v2 chan int
	var v22 v2
	Print(`kind is:`, reflect.TypeOf(v22).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(v22).Name() == ``)

	i1 = func() {}
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)
	type v3 func()
	var v33 v3
	Print(`kind is:`, reflect.TypeOf(v33).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(v33).Name() == ``)

	i1 = (*interface{})(nil)
	Print(`kind is:`, reflect.TypeOf(i1).Elem().Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Elem().Name() == ``)
	i1 = (*error)(nil)
	Print(`kind is:`, reflect.TypeOf(i1).Elem().Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Elem().Name() == ``)

	i1 = map[string]string{}
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)
	type v4 map[string]string
	var v44 v4
	Print(`kind is:`, reflect.TypeOf(v44).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(v44).Name() == ``)

	i1 = new(int)
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)
	type v5 *int
	var v55 v5
	Print(`kind is:`, reflect.TypeOf(v55).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(v55).Name() == ``)

	i1 = []int{1, 2, 3}
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)
	type v6 []int
	var v66 v6
	Print(`kind is:`, reflect.TypeOf(v66).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(v66).Name() == ``)

	i1 = os.File{}
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)
	i1 = struct{}{}
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)

	i1 = unsafe.Pointer(new(int))
	Print(`kind is:`, reflect.TypeOf(i1).Kind().String()+`,`, `unnamed:`, reflect.TypeOf(i1).Name() == ``)
}

func main() {
	Case1()
}
