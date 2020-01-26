package main

import (
	"github.com/qingsong-he/ce"
	"os"
)

func init() {
	ce.Print(os.Args[0])
}

func main() {
	var i uint8 = 0xff

	// clear 4th bit
	i &^= 8
	ce.Printf("%08b\n", i)

	// set 4th bit
	i |= 8
	ce.Printf("%08b\n", i)
}
