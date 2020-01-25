package main

import (
	"github.com/monnand/dhkx"
	"github.com/qingsong-he/ce"
	"math/big"
	"os"
)

func init() {
	ce.Print(os.Args[0])
}

func main() {
	{
		// common
		g := big.NewInt(3)
		p := big.NewInt(0).Sub(
			big.NewInt(0).Exp(big.NewInt(2), big.NewInt(127), nil),
			big.NewInt(1),
		)

		// a
		a := big.NewInt(100)
		ga := big.NewInt(0).Exp(
			g,
			a,
			p,
		)

		// b
		b := big.NewInt(101)
		gb := big.NewInt(0).Exp(
			g,
			b,
			p,
		)
		ce.Print(ga, gb)

		//
		// exchange:
		//

		// a
		gb_a := big.NewInt(0).Exp(
			gb,
			a,
			p,
		)

		ce.Print(gb_a)

		// b
		ga_b := big.NewInt(0).Exp(
			ga,
			b,
			p,
		)
		ce.Print(ga_b)
	}

	{
		// common
		g := big.NewInt(3)
		p := big.NewInt(0).Sub(
			big.NewInt(0).Exp(big.NewInt(2), big.NewInt(127), nil),
			big.NewInt(1),
		)

		group := dhkx.CreateGroup(p, g)
		a, err := group.GeneratePrivateKey(nil)
		ce.CheckError(err)

		b, err := group.GeneratePrivateKey(nil)
		ce.CheckError(err)
		ce.Print(a.String(), b.String())

		//
		// exchange:
		//

		// a
		aResult, err := group.ComputeKey(b, a)
		ce.CheckError(err)
		ce.Print(aResult.String())

		// b
		bResult, err := group.ComputeKey(a, b)
		ce.CheckError(err)
		ce.Print(bResult.String())
	}

}
