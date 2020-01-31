package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"github.com/qingsong-he/ce"
	"os"
	"reflect"
)

func init() {
	ce.Print(os.Args[0])
}

func case1() {

	h256 := sha256.New()

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	ce.CheckError(err)

	msg := "hello"

	// enc with public key
	msgByEnc, err := rsa.EncryptOAEP(h256, rand.Reader, &privKey.PublicKey, []byte(msg), nil)
	ce.CheckError(err)

	// dec with private key
	msgByDec, err := rsa.DecryptOAEP(h256, rand.Reader, privKey, msgByEnc, nil)
	ce.CheckError(err)
	ce.Print(reflect.DeepEqual([]byte(msg), msgByDec))

	// sign with private key, and verify with public key
	msgForSign := []byte("i am hello")
	h256.Reset()
	_, err = h256.Write(msgForSign)
	ce.CheckError(err)

	msgForSignByHash := h256.Sum(nil)
	msgForSignBySign, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, msgForSignByHash)
	ce.CheckError(err)

	// damage this verify
	if false {
		msgForSignBySign = append(msgForSignBySign, byte('a'))
	}

	err = rsa.VerifyPKCS1v15(&privKey.PublicKey, crypto.SHA256, msgForSignByHash, msgForSignBySign)
	ce.Print(err == nil)
	ce.CheckError(err)
}

func main() {
	case1()
}
