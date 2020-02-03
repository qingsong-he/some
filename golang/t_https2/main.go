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
	// openssl genrsa -out client.key 2048
	// openssl req -new -key client.key -subj "/CN=client" -out client.csr
	// openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000
	go func() {
		caCrtBin, err := ioutil.ReadFile("ca.crt")
		ce.CheckError(err)
		crtPool := x509.NewCertPool()
		crtPool.AppendCertsFromPEM(caCrtBin)

		s := http.Server{
			Addr:    ":3000",
			Handler: http.DefaultServeMux,
			TLSConfig: &tls.Config{
				ClientAuth: tls.RequireAndVerifyClientCert,
				ClientCAs:  crtPool,
			},
		}
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("i am root"))
		})
		err = s.ListenAndServeTLS("server.crt", "server.key")
		ce.CheckError(err)
	}()

	// client 1
	{
		caCrtBin, err := ioutil.ReadFile("ca.crt")
		ce.CheckError(err)
		crtPool := x509.NewCertPool()
		crtPool.AppendCertsFromPEM(caCrtBin)

		cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
		ce.CheckError(err)

		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{RootCAs: crtPool, Certificates: []tls.Certificate{cliCrt}}
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
