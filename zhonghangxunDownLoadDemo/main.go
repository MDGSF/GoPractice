package main

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/MDGSF/utils/log"
)

func main() {
	var err error
	err = errors.New("my error")
	str := fmt.Sprintf("%v", err)
	fmt.Println(str)
}

func main1() {
	//CAPath := "ca-cert.pem"
	CertPath := "client-cert.pem"
	KeyPath := "client-key.pem"
	var err error

	//rootPEM, err := ioutil.ReadFile(CAPath)
	//if err != nil {
	//	log.Error("read file %v failed, err = %v", CAPath, err)
	//	return
	//}

	//roots := x509.NewCertPool()
	//ok := roots.AppendCertsFromPEM(rootPEM)
	//if !ok {
	//	log.Error("AppendCertsFromPEM failed")
	//	return
	//}

	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true //just for test.
	//tlsConfig.ServerName = "117.34.118.23"
	//tlsConfig.RootCAs = roots
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(CertPath, KeyPath)

	tr := &http.Transport{TLSClientConfig: tlsConfig}

	client := &http.Client{Transport: tr}

	uri := "https://117.34.118.23:5632/upload/UpgradePak/20180521154901_4.gz"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	status := response.StatusCode
	if status != 200 {
		log.Error("status code = %v", status)
	}
	defer response.Body.Close()

	saveFileName := "20180521154901_4.gz"

	saveFile, err := os.OpenFile(saveFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Error("open save file %v failed, err = %v", saveFileName, err)
		return
	}
	defer saveFile.Close()

	r := bufio.NewReader(response.Body)
	w := bufio.NewWriter(saveFile)

	count := 0
	buf := make([]byte, 4096)
	for {
		rn, rerr := r.Read(buf)
		wn, werr := w.Write(buf[:rn])
		count += wn

		if rerr != nil {
			log.Printf("read failed, rn = %v, wn = %v, err = %v\n", rn, wn, rerr)
			break
		}
		if werr != nil {
			log.Error("write failed, werr = %v", werr)
			break
		}
		if rn != wn {
			log.Error("rn = %v, wn = %v", rn, wn)
		}
	}
	w.Flush() //flush all buffered data into file.
	fmt.Println("count = ", count)

	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Error("read body failed, err = %v", err)
	//	panic(err)
	//}

	//fmt.Println("len body = ", len(body))

	//_, err = saveFile.Write(body)
	//if err != nil {
	//	log.Error("save file %v write failed, err = %v", saveFileName, err)
	//	return
	//}
}

//func (c *TS0HTTPClient) genHTTPSClient() *http.Client {
//	rootPEM, err := ioutil.ReadFile(MainConfig.CAPath)
//	if err != nil {
//		log.Error("read file %v failed, err = %v", MainConfig.CAPath, err)
//		return nil
//	}
//
//	roots := x509.NewCertPool()
//	ok := roots.AppendCertsFromPEM(rootPEM)
//	if !ok {
//		log.Error("AppendCertsFromPEM failed")
//		return nil
//	}
//
//	tlsConfig := &tls.Config{}
//	tlsConfig.ServerName = MainConfig.S0Addr.Host
//	tlsConfig.RootCAs = roots
//	tlsConfig.Certificates = make([]tls.Certificate, 1)
//	tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(MainConfig.CertPath, MainConfig.KeyPath)
//
//	tr := &http.Transport{TLSClientConfig: tlsConfig}
//
//	client := &http.Client{Transport: tr}
//	return client
//}
//
