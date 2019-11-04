package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new(cgi.Handler)
		handler.Path = "/home/huangjian/a/local/go/bin/go"

		scriptDir := "/home/huangjian/a/local/gopath/src/github.com/MDGSF/GoPractice/cgiDemo/1"
		script := scriptDir + r.URL.Path
		handler.Dir = scriptDir
		args := []string{"run", script}
		handler.Args = append(handler.Args, args...)
		handler.Env = append(handler.Env, "HOME=/home/huangjian")
		handler.Env = append(handler.Env, "GOPATH=/home/huangjian/a/local/gopath")
		handler.Env = append(handler.Env, "GOROOT=/home/huangjian/a/local/go")

		log.Println(script)
		log.Println(handler.Args)

		handler.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
