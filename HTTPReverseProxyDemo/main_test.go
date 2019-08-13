package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MDGSF/utils/log"
)

func TestMux(t *testing.T) {
	initMuxTestable()

	server := httptest.NewServer(muxTestable)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		log.Fatalf("URL = %v, err = %v", server.URL, err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	expected := "I'm muxTestable."
	if string(content) != expected {
		log.Fatalf("content = %v, expected = %v", content, expected)
	}
}
