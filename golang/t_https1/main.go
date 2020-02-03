package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/qingsong-he/ce"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	ce.Print(os.Args[0])
}

func case1() {
	// openssl genrsa -out ca.key 2048
	// openssl req -x509 -new -key ca.key -subj "/CN=foo.bar" -days 5000 -out ca.crt
	// openssl genrsa -out server.key 2048
	// openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
	// openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("i am root"))
		})
		err := http.ListenAndServeTLS(":3000", "server.crt", "server.key", nil)
		ce.CheckError(err)
	}()

	// client 1
	{
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		resp, err := http.Get("https://localhost:3000")
		ce.CheckError(err)
		body, err := ioutil.ReadAll(resp.Body)
		ce.CheckError(err)
		defer resp.Body.Close()
		ce.Print(string(body))
	}

	// client 2
	{
		caCrtBin, err := ioutil.ReadFile("ca.crt")
		ce.CheckError(err)
		crtPool := x509.NewCertPool()
		crtPool.AppendCertsFromPEM(caCrtBin)
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{RootCAs: crtPool}
		resp, err := http.Get("https://localhost:3000")
		ce.CheckError(err)
		body, err := ioutil.ReadAll(resp.Body)
		ce.CheckError(err)
		defer resp.Body.Close()
		ce.Print(string(body))
	}
}

func main() {
	case1()
}
