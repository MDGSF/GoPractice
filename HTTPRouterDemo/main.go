package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Auth(h httprouter.Handle, role, privilage string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Printf("role = %v, privilage = %v\n", role, privilage)
		h(w, r, ps)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Index\n")
}

func Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Protected\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/protected/", Auth(Protected, "root", "uploadPackage"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
