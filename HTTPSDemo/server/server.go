package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("https demo server."))
}

var CAPath = "/usr/local/gopath/src/github.com/MDGSF/GoPractice/HTTPSDemo/ca.cer"
var ServerCertPath = "/usr/local/gopath/src/github.com/MDGSF/GoPractice/HTTPSDemo/server.cer"
var ServerKeyPath = "/usr/local/gopath/src/github.com/MDGSF/GoPractice/HTTPSDemo/server-key-out.pem"

func main() {
	rootPEM, err := ioutil.ReadFile(CAPath)
	if err != nil {
		log.Fatal("read ca.cer failed, err = %v", err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		log.Fatal("roots append failed")
	}

	tlsConfig := &tls.Config{}
	tlsConfig.ClientCAs = roots
	tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert

	s := &http.Server{}
	s.Addr = "localhost:10000"
	s.Handler = &myHandler{}
	s.TLSConfig = tlsConfig

	err = s.ListenAndServeTLS(ServerCertPath, ServerKeyPath)
	if err != nil {
		log.Fatalf("ListenAndServeTLS err = %v", err)
	}
}
