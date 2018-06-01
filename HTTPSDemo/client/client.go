package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var CAPath = "/usr/local/gopath/src/github.com/MDGSF/GoPractice/HTTPSDemo/ca.cer"
var ClientCertPath = "/usr/local/gopath/src/github.com/MDGSF/GoPractice/HTTPSDemo/client.cer"
var ClientKeyPath = "/usr/local/gopath/src/github.com/MDGSF/GoPractice/HTTPSDemo/client-key-out.pem"

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
	//tlsConfig.ServerName = "localhost"
	tlsConfig.RootCAs = roots
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(ClientCertPath, ClientKeyPath)
	if err != nil {
		log.Fatal("load x509 failed, err = %v", err)
	}

	tr := &http.Transport{}
	tr.TLSClientConfig = tlsConfig

	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:10000")
	if err != nil {
		log.Fatal("client get failed, err = %v", err)
	}

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(strings.TrimSpace(string(content)))
}
