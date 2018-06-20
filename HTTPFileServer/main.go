package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MDGSF/utils/log"
)

func main() {
	fmt.Printf("%v", 32<<20)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe("192.168.1.178:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {

	log.Info("r = %v", r)
	log.Info("r.Method = %v", r.Method)
	log.Info("r.URL = %v", r.URL)
	log.Info("r.Header = %v", r.Header)
	log.Info("r.ContentLength = %v", r.ContentLength)
	log.Info("r.Host = %v", r.Host)
	log.Info("r.Form = %v", r.Form)
	log.Info("r.PostForm = %v", r.PostForm)
	log.Info("r.RemoteAddr = %v", r.RemoteAddr)
	log.Info("r.RequestURI = %v", r.RequestURI)

	switch r.Method {
	case "PUT":
		uploadPut(w, r)
	case "GET":
		uploadGet(w, r)
	case "DELETE":
		uploadDel(w, r)
	}
}

func uploadPut(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Error("%v", err)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error("%v", err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
}

func uploadGet(w http.ResponseWriter, r *http.Request) {
}

func uploadDel(w http.ResponseWriter, r *http.Request) {
}
