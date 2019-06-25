package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/astaxie/beego/httplib"
)

func main() {
	url := "http://127.0.0.1:12345/ft/uploads/"
	sendHead(url)
	sendPut(url)
	sendHead(url)
	sendGet(url)
	sendDel(url)
	sendHead(url)
}

func sendHead(url string) {
	req := httplib.Head(url + calculateFileSha1("test.txt"))
	req.Param("user_id", "testusername")
	req.Param("user_key", "testuserkey")
	req.SetTimeout(5*time.Second, 5*time.Second)
	rep, err := req.Response()
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	log.Info("rep.StatusCode = %v", rep.StatusCode)
}

func sendPut(url string) {
	req := httplib.Put(url + calculateFileSha1("test.txt"))
	req.Param("user_id", "testusername")
	req.Param("user_key", "testuserkey")
	req.PostFile("file", "test.txt")
	req.SetTimeout(5*time.Second, 5*time.Second)
	rep, err := req.Response()
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	log.Info("rep.StatusCode = %v", rep.StatusCode)
}

func sendGet(url string) {
	req := httplib.Get(url + calculateFileSha1("test.txt"))
	req.Param("user_id", "testusername")
	req.Param("user_key", "testuserkey")
	req.SetTimeout(5*time.Second, 5*time.Second)
	rep, err := req.Response()
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	log.Info("rep.StatusCode = %v", rep.StatusCode)

	if err := req.ToFile("newtest.txt"); err != nil {
		log.Error("err = %v", err)
		return
	}
}

func sendDel(url string) {
	req := httplib.Delete(url + calculateFileSha1("test.txt"))
	req.Param("user_id", "testusername")
	req.Param("user_key", "testuserkey")
	req.SetTimeout(5*time.Second, 5*time.Second)
	rep, err := req.Response()
	if err != nil {
		log.Error("err = %v", err)
		return
	}
	log.Info("rep.StatusCode = %v", rep.StatusCode)
}

func calculateFileSha1(strFileName string) string {
	aucData, _ := ioutil.ReadFile(strFileName)
	hasher := sha1.New()
	hasher.Write(aucData)
	return hex.EncodeToString(hasher.Sum(nil))
}
