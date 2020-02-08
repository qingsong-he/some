package main

import (
	"github.com/qingsong-he/ce"
	"os"
	"unicode"
	"unicode/utf8"
)

func init() {
	ce.Print(os.Args[0])
}

func case1() {
	var all string

	all = "世界"
	ce.Print(all)

	all = "\xe4\xb8\x96\xe7\x95\x8c"
	ce.Print(all)

	all = "\u4e16\u754c"
	ce.Print(all)

	all = "\U00004e16\U0000754c"
	ce.Print(all)

	var c rune

	c = '世'
	ce.Printf("%c", c)

	c = '\u4e16'
	ce.Printf("%c", c)

	c = '\U00004e16'
	ce.Printf("%c", c)
}

func case2() {
	s := "Hello，世界"
	ce.Print(len(s))
	ce.Print(utf8.RuneCountInString(s))
	for i, r := range s {
		if unicode.Is(unicode.Han, r) {
			ce.Printf("%d Han %c", i, r)
		} else {
			ce.Printf("%d %c", i, r)
		}
	}

	// '\uFFFD' maybe indicate an error

	// string -> []rune
	// []rune -> string

	// string -> []byte
	// []byte -> string

}

func main() {
	case2()
}
