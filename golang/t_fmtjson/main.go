package main

import (
	"encoding/json"
	"github.com/qingsong-he/ce"
	"os"
)

func init() {
	ce.Print(os.Args[0])
}

func fmtJsonStr(jsonByNotFmt, indent string) (string, error) {
	obj := new(interface{})
	err := json.Unmarshal([]byte(jsonByNotFmt), obj)
	if err != nil {
		return "", err
	}
	json1, err := json.MarshalIndent(obj, "", indent)
	return string(json1), err
}

func main() {
	json1 := `{"hello": 1.1}`
	json1ByFmt, err := fmtJsonStr(json1, " ")
	ce.CheckError(err)
	println(json1ByFmt)
}
