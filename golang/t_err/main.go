package main

import (
	"errors"
	"fmt"
	"github.com/qingsong-he/ce"
	"os"
)

func init() {
	ce.Print(os.Args[0])
}

func main() {

	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)

	ce.Print(err3)                               // err3: [err2: [new error]]
	ce.Print(errors.Unwrap(err3))                // err2: [new error]
	ce.Print(errors.Unwrap(errors.Unwrap(err3))) // new error
	ce.Print(errors.Unwrap(nil))                 // nil
	ce.Print(errors.Is(err3, err2))              // true
	ce.Print(errors.Is(err3, err1))              // true

	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			ce.Print("Failed at path:", pathError.Path)
		} else {
			ce.Print(err)
		}
	}
}
