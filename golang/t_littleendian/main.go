package main

import (
	"bytes"
	"encoding/binary"
	. "github.com/qingsong-he/ce"
)

func IsLittleEndian() bool {
	var i uint16 = 0x0102
	bytesBuf := bytes.NewBuffer(nil)
	err := binary.Write(bytesBuf, binary.LittleEndian, i)
	if err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint16(bytesBuf.Bytes()) == i
}

func Case1() {

	// is little endian?
	Print(IsLittleEndian())

	// convert number to byte slice by little endian
	var i uint16 = 0x0102
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, i)
	CheckError(err)
	iByByteSlic := buf.Bytes()
	Print(iByByteSlic)

	// convert number to byte slice by big endian
	buf1 := bytes.NewBuffer(nil)
	err1 := binary.Write(buf1, binary.BigEndian, i)
	CheckError(err1)
	iByByteSlic1 := buf1.Bytes()
	Print(iByByteSlic1)

	// convert byte slice to number by little and big endian
	Print(binary.LittleEndian.Uint16(iByByteSlic), binary.BigEndian.Uint16(iByByteSlic))

	// convert number to big endian byte slic
	iByBigEndianSlic := make([]byte, 2)
	binary.BigEndian.PutUint16(iByBigEndianSlic, i)
	Print(iByBigEndianSlic)

	// convert big endian byte slic to little endian number
	iByLittleEndianSlic := make([]byte, 2)
	binary.LittleEndian.PutUint16(iByLittleEndianSlic, binary.BigEndian.Uint16(iByBigEndianSlic))
	Print(iByLittleEndianSlic, binary.LittleEndian.Uint16(iByLittleEndianSlic))
}

func main() {
	Case1()
}
