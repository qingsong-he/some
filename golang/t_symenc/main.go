package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	. "github.com/qingsong-he/ce"
	"io"
)

func Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	return append(src, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func UnPadding(src []byte) []byte {
	length := len(src)
	return src[:(length - int(src[length-1]))]
}

func DESEncByECB(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	src = Padding(src, bs)
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func DESDecByECB(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return UnPadding(out), nil
}

func DESEncByCBC(src, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	src = Padding(src, block.BlockSize())
	out := make([]byte, len(src))

	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(out, src)
	return out, nil
}

func DESDecByCBC(src, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	bm := cipher.NewCBCDecrypter(block, iv)
	out := make([]byte, len(src))
	bm.CryptBlocks(out, src)
	return UnPadding(out), nil
}

func AESEncByCFB(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	out := make([]byte, aes.BlockSize+len(src))
	iv := out[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cipher.NewCFBEncrypter(block, iv).XORKeyStream(out[aes.BlockSize:], src)
	return out, nil
}

func AESDecByCFB(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := src[:aes.BlockSize]
	out := src[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(out, out)
	return out, nil
}

func main() {
	keyBy8Byte := []byte("12345678")
	ivBy8Byte := []byte("87654321")

	keyBy16Byte := []byte("1234567812345678")

	// test des ecb
	if true {
		result, err := DESEncByECB([]byte("hello world"), keyBy8Byte)
		CheckError(err)
		Print(result)

		result1, err := DESDecByECB(result, keyBy8Byte)
		CheckError(err)
		Print(string(result1))
	}

	// test des cbc
	if true {
		result, err := DESEncByCBC([]byte("hello world"), keyBy8Byte, ivBy8Byte)
		CheckError(err)
		Print(result)

		result1, err := DESDecByCBC(result, keyBy8Byte, ivBy8Byte)
		CheckError(err)
		Print(string(result1))
	}

	// test aes cfb
	if true {
		result, err := AESEncByCFB([]byte("hello world"), keyBy16Byte)
		CheckError(err)
		Print(result)

		result1, err := AESDecByCFB(result, keyBy16Byte)
		CheckError(err)
		Print(string(result1))
	}
}
