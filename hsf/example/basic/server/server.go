package main

import (
	"net/http"

	"github.com/MDGSF/hsf"
)

func main() {

	m := make(map[string]*hsf.TFileInfo)

	IsValidUser := func(w http.ResponseWriter, r *http.Request) bool {
		return true
	}

	SetUploadedFileInfo := func(strFileKey string, pstFileInfo *hsf.TFileInfo) {
		m[strFileKey] = pstFileInfo
	}

	GetUploadedFileInfo := func(strFileKey string) *hsf.TFileInfo {
		if pstFileInfo, ok := m[strFileKey]; ok {
			return pstFileInfo
		}
		return nil
	}

	DelUploadedFileInfo := func(strFileKey string) {
		delete(m, strFileKey)
	}

	pstServer := hsf.NewHTTPStaticFileServer(
		"data/uploads",
		IsValidUser,
		SetUploadedFileInfo,
		GetUploadedFileInfo,
		DelUploadedFileInfo,
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/ft/uploads/", pstServer.Upload)

	httpServer := &http.Server{}
	httpServer.Addr = "127.0.0.1:12345"
	httpServer.Handler = mux

	httpServer.ListenAndServe()
}
