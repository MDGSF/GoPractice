package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/MDGSF/utils/log"
)

const DataPath = "./data/"

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ft/uploads/", upload)
	err := http.ListenAndServe("192.168.1.178:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Info("index start")
}

func upload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	fmt.Println()
	log.Info("--------------------------------------------")
	log.Info("r = %v", r)
	log.Info("r.Method = %v", r.Method)
	log.Info("r.URL = %v", r.URL)

	log.Info("r.URL.Scheme = %v", r.URL.Scheme)
	log.Info("r.URL.Opaque = %v", r.URL.Opaque)
	log.Info("r.URL.User = %v", r.URL.User)
	log.Info("r.URL.Host = %v", r.URL.Host)
	log.Info("r.URL.Path = %v", r.URL.Path)

	log.Info("r.URL.Path last = %v", filepath.Base(r.URL.Path))

	log.Info("r.URL.RawPath = %v", r.URL.RawPath)
	log.Info("r.URL.ForceQuery = %v", r.URL.ForceQuery)
	log.Info("r.URL.RawQuery = %v", r.URL.RawQuery)
	log.Info("r.URL.Fragment = %v", r.URL.Fragment)

	log.Info("r.Header = %v", r.Header)
	log.Info("r.ContentLength = %v", r.ContentLength)
	log.Info("r.Host = %v", r.Host)
	log.Info("r.Form = %v", r.Form)
	log.Info("r.PostForm = %v", r.PostForm)
	log.Info("r.RemoteAddr = %v", r.RemoteAddr)
	log.Info("r.RequestURI = %v", r.RequestURI)
	log.Info("--------------------------------------------")

	switch r.Method {
	case "PUT":
		uploadPut(w, r)
	case "GET":
		uploadGet(w, r)
	case "DELETE":
		uploadDel(w, r)
	default:
		log.Error("Unknown method = %v", r.Method)
	}
}

func calcDigest(r io.Reader) string {
	h := sha1.New()
	io.Copy(h, r)
	return hex.EncodeToString(h.Sum(nil))
}

func uploadPut(w http.ResponseWriter, r *http.Request) {
	log.Info("uploadPut start")
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Error("%v", err)
		return
	}
	defer file.Close()

	log.Info("sha1 = %v", calcDigest(file))
	file.Seek(io.SeekStart, 0)

	log.Info("handler.Filename = %v", handler.Filename)
	newFileName := filepath.Base(r.URL.Path)
	newFileName = DataPath + newFileName

	f, err := os.OpenFile(newFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Error("%v", err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	w.Write([]byte("uploadPut"))
}

func uploadGet(w http.ResponseWriter, r *http.Request) {
	log.Info("uploadGet start")

	fileName := filepath.Base(r.URL.Path)
	fileName = DataPath + fileName
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Error("%v", err)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func uploadDel(w http.ResponseWriter, r *http.Request) {
	log.Info("uploadDel start")
	fileName := filepath.Base(r.URL.Path)
	os.Remove(DataPath + fileName)
	w.Write([]byte("uploadDel"))
}
