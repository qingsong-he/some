package main

import (
	"bytes"
	"fmt"
	. "github.com/qingsong-he/ce"
	"io"
)

type t1 int

func (t *t1) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (t *t1) Write(p []byte) (n int, err error) {
	return 0, nil
}

func Case1() {
	var i1 io.ReadWriter
	var i2 io.ReadWriter
	i1 = bytes.NewBuffer(nil)
	i2 = new(t1)

	{
		v, ok := i1.(*bytes.Buffer)
		Print(v, ok) // true
	}

	{
		v, ok := i2.(*t1)
		Print(v, ok) // true
	}

	{
		v, ok := i1.(*t1)
		Print(v, ok) // false
	}

	{
		v, ok := i2.(*bytes.Buffer)
		Print(v, ok) // false
	}
}

func Case2() {
	var i1 io.ReadWriter
	var i2 io.Writer

	b1 := bytes.NewBuffer(nil)
	i1 = b1
	i2 = b1

	// 'i1 = i2' can not compile
	i2 = i1

	Print(i1, i2)

	// runtime query
	v, ok := i2.(io.ReadWriter)
	Print(v, ok)
}

type foobarr interface {
	func1()
	func2()
}

type foobar struct {
	some string
}

func (c foobar) func1() {
	fmt.Println("func1:", c.some)
}

func (c *foobar) func2() {
	fmt.Println("func2:", c.some)
}

func Case3() {
	c := foobar{some: "foo"}
	defer c.func1() // foo

	c.some = "bar"
	Print(c.some) // bar
}

func Case4() {
	var fbr foobarr
	var fb1 foobar
	fb1.some = "foobar"

	// 'fbr = fb1' can not compile
	fbr = &fb1

	fbr.func1()
	fbr.func2()
}

func Case5() {
	var i1 interface{}

	func1 := func() {
		switch t1 := i1.(type) {
		default:
			Print(t1)
		}
	}

	func1() // nil

	i1 = 1
	func1() // int

	i1 = 1.1
	func1() // float64
}

func main() {
	Case1()
}
