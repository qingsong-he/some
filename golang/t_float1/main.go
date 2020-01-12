package main

import (
	. "github.com/qingsong-he/ce"
	"math"
)

func main() {
	var f1 = math.MaxFloat64
	var f11 = math.SmallestNonzeroFloat64
	var f2 = math.Inf(1)
	var f3 = math.Inf(-1)
	var f4 = math.NaN()

	Printf("%064b", math.Float64bits(f1))
	Printf("%064b", math.Float64bits(f11))
	Printf("%064b", math.Float64bits(f2))
	Printf("%064b", math.Float64bits(f3))
	Printf("%064b", math.Float64bits(f4))
}
