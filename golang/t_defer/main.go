package main

import (
	. "github.com/qingsong-he/ce"
)

func Case1() {
	i := 1
	defer Print(i) // 1, but log file location is not righlog.
	i++
	Print(i) // 2
}

func Case2() {
	i := 1
	defer func() {
		Print(i) // 2
	}()
	i++
	Print(i) // 2
}

type foobar struct {
	some string
}

func (c foobar) func1() {
	Print("func1:", c.some)
}

func (c *foobar) func2() {
	Print("func2:", c.some)
}

func Case3() {
	c := foobar{some: "foo"}
	defer c.func1() // foo

	c.some = "bar"
	Print(c.some) // bar
}

func Case4() {
	c := foobar{some: "foo"}
	defer func() {
		c.func1() // bar
	}()

	c.some = "bar"
	Print(c.some) // bar
}

func Case5() {
	c := foobar{some: "foo"}
	defer c.func2() // bar

	c.some = "bar"
	Print(c.some) // bar
}

func Case6() {
	func1 := func() (r int) {
		defer func() {
			r++
		}()
		return 0
	}

	Print(func1())
}

func main() {
	Case1()
}
