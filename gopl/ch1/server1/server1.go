//Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.URL.Path = %q\n", r.URL.Path)
}
