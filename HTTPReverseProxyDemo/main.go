package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/MDGSF/utils/log"
)

func startSimpleHTTPService() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8887", nil))
}

func handleDefault(w http.ResponseWriter, req *http.Request) {
	log.Printf("handleDefault, %v %v", req.Method, req.URL.Path)
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func startService() {
	http.HandleFunc("/", handleDefault)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

/*
proxy localhost:8889 request to localhost:8888
*/
func startReverseProxy() {
	go startService()

	rpURL, err := url.Parse("http://localhost:8888")
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	apiReverseProxy := httputil.NewSingleHostReverseProxy(rpURL)
	http.Handle("/api/", http.StripPrefix("/api/", apiReverseProxy))
	log.Fatal(http.ListenAndServe(":8889", nil))
}

func startMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("mux handle, %v %v", req.Method, req.URL.Path)
		w.WriteHeader(200)
		w.Write([]byte("I'm mux."))
	})
	log.Fatal(http.ListenAndServe("localhost:8890", mux))
}

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("mux handle, %v %v", req.Method, req.URL.Path)
		w.WriteHeader(200)
		w.Write([]byte("I'm http server mux."))
	})

	server := &http.Server{
		Addr:    "localhost:8891",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}

var muxTestable *http.ServeMux

func initMuxTestable() {
	muxTestable = http.NewServeMux()
	muxTestable.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("muxTestable handle, %v %v", req.Method, req.URL.Path)
		w.WriteHeader(200)
		w.Write([]byte("I'm muxTestable."))
	})
}

func startMuxTestable() {
	initMuxTestable()
	log.Fatal(http.ListenAndServe("localhost:8892", muxTestable))
}

func main() {
	go startSimpleHTTPService()
	go startReverseProxy()
	go startMux()
	go startServer()
	go startMuxTestable()

	shutdown := make(chan struct{})
	<-shutdown
}
