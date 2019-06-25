package hsf

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/MDGSF/utils/log"
)

// TFileInfo file info
type TFileInfo struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Time  int64  `json:"time"`
}

// THTTPStaticFileServer HTTP static file server.
type THTTPStaticFileServer struct {
	// DataDir use to store upload file, default is ./data/uploads
	DataDir string

	// IsValidUser check whether request user is valid or not.
	IsValidUser func(w http.ResponseWriter, r *http.Request) bool

	// SetUploadedFileInfo store file info
	SetUploadedFileInfo func(strFileKey string, pstFileInfo *TFileInfo)

	// GetUploadedFileInfo get file info
	GetUploadedFileInfo func(strFileKey string) *TFileInfo

	// DelUploadedFileInfo delete file info
	DelUploadedFileInfo func(strFileKey string)
}

// NewHTTPStaticFileServer create new  HTTP static file server.
func NewHTTPStaticFileServer(
	DataDir string,
	IsValidUser func(w http.ResponseWriter, r *http.Request) bool,
	SetUploadedFileInfo func(strFileKey string, pstFileInfo *TFileInfo),
	GetUploadedFileInfo func(strFileKey string) *TFileInfo,
	DelUploadedFileInfo func(strFileKey string),
) *THTTPStaticFileServer {

	if DataDir == "" ||
		IsValidUser == nil ||
		SetUploadedFileInfo == nil ||
		GetUploadedFileInfo == nil ||
		DelUploadedFileInfo == nil {
		log.Error("Invalid parameters.")
		return nil
	}

	pstServer := &THTTPStaticFileServer{
		DataDir:             DataDir,
		IsValidUser:         IsValidUser,
		SetUploadedFileInfo: SetUploadedFileInfo,
		GetUploadedFileInfo: GetUploadedFileInfo,
		DelUploadedFileInfo: DelUploadedFileInfo,
	}

	TryMkDir(DataDir)

	return pstServer
}

// Upload process HTTP upload request.
func (pstServer *THTTPStaticFileServer) Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		// log.Info("%v, try to ParseForm...\n", err)
		if err := r.ParseForm(); err != nil {
			log.Error("parse form failed, err = %v\n", err)
		}
	}

	if pstServer.IsValidUser != nil {
		if !pstServer.IsValidUser(w, r) {
			HTTPRsp(w, http.StatusForbidden, "invalid id or key")
			return
		}
	}

	log.Info("[%v](%v)", r.Method, r.URL.Path)

	switch r.Method {
	case "HEAD":
		pstServer.uploadHead(w, r)
	case "PUT":
		pstServer.uploadPut(w, r)
	case "GET":
		pstServer.uploadGet(w, r)
	case "DELETE":
		pstServer.uploadDelete(w, r)
	default:
		log.Error("Unknown method = %v", r.Method)
	}
}

func (pstServer *THTTPStaticFileServer) uploadHead(w http.ResponseWriter, r *http.Request) {
	strFileKey := getFileKeyFromURLPath(r.URL.Path)
	pstFileInfo := pstServer.GetUploadedFileInfo(strFileKey)
	if pstFileInfo == nil {
		log.Error("strFileKey = %v, is not exists.", strFileKey)
		HTTPRsp(w, http.StatusBadRequest, "get request uploadHead failed")
		return
	}

	log.Info("get request uploadHead success, strFileKey = %v, pstFileInfo = %v", strFileKey, pstFileInfo)

	w.Header().Set("x-file-owner", pstFileInfo.Owner)
	w.Header().Set("x-file-name", pstFileInfo.Name)
	w.Header().Set("x-file-size", strconv.FormatInt(pstFileInfo.Size, 10))
	w.Header().Set("x-file-utime", strconv.FormatInt(pstFileInfo.Time, 10))
	w.WriteHeader(http.StatusOK)
}

func (pstServer *THTTPStaticFileServer) uploadPut(w http.ResponseWriter, r *http.Request) {
	stFile, pstFileHeader, err := r.FormFile("file")
	if err != nil {
		log.Error("%v", err)
		HTTPRsp(w, http.StatusBadRequest, "")
		return
	}
	defer stFile.Close()

	strFileKey1 := calculateFileKey(stFile)
	strFileKey2 := getFileKeyFromURLPath(r.URL.Path)

	if strFileKey1 != strFileKey2 {
		log.Error("strFileKey1 = %v, strFileKey2 = %v", strFileKey1, strFileKey2)
		HTTPRsp(w, http.StatusBadRequest, "")
		return
	}

	TryMkDir(pstServer.genFileDir(strFileKey1))

	strNewFileName := pstServer.genFilePath(strFileKey1)
	pstNewFile, err := os.OpenFile(strNewFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Error("open strNewFileName = %v failed, err = %v", strNewFileName, err)
		HTTPRsp(w, http.StatusBadRequest, "")
		return
	}
	defer pstNewFile.Close()

	io.Copy(pstNewFile, stFile)

	pstFileInfo := &TFileInfo{}
	pstFileInfo.Owner = r.Form.Get("user_id")
	pstFileInfo.Name = pstFileHeader.Filename
	pstFileInfo.Size = pstFileHeader.Size
	pstFileInfo.Time = time.Now().UnixNano()
	pstServer.SetUploadedFileInfo(strFileKey1, pstFileInfo)

	type TResponse struct {
		Result string      `json:"result"`
		Data   interface{} `json:"data"`
	}
	rsp := &TResponse{Result: "ok", Data: pstFileInfo}
	b, _ := json.Marshal(rsp)
	w.Write(b)
}

func (pstServer *THTTPStaticFileServer) uploadGet(w http.ResponseWriter, r *http.Request) {
	strFilePath := pstServer.genFilePath(getFileKeyFromURLPath(r.URL.Path))
	pstFile, err := os.OpenFile(strFilePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Error("open file %v failed, err = %v", strFilePath, err)
		HTTPRsp(w, http.StatusBadRequest, "")
		return
	}
	defer pstFile.Close()
	io.Copy(w, pstFile)
}

func (pstServer *THTTPStaticFileServer) uploadDelete(w http.ResponseWriter, r *http.Request) {
	pstServer.DelUploadedFileInfo(getFileKeyFromURLPath(r.URL.Path))
	os.Remove(pstServer.genFilePath(getFileKeyFromURLPath(r.URL.Path)))
	w.WriteHeader(http.StatusNoContent)
}

// genFilePath generate file absolute path.
func (pstServer *THTTPStaticFileServer) genFilePath(strFileKey string) string {
	return filepath.Join(pstServer.DataDir, strFileKey[:2], strFileKey)
}

// genFileDir generate file parent directory.
func (pstServer *THTTPStaticFileServer) genFileDir(strFileKey string) string {
	return filepath.Join(pstServer.DataDir, strFileKey[:2])
}

// getFileKeyFromURLPath extract file key from URL path.
func getFileKeyFromURLPath(strURLPath string) string {
	return filepath.Base(strURLPath)
}

// calculateFileKey calculate file key
func calculateFileKey(stFile multipart.File) string {
	h := sha1.New()
	io.Copy(h, stFile)
	strFileKey := hex.EncodeToString(h.Sum(nil))
	stFile.Seek(io.SeekStart, 0)
	return strFileKey
}

/*
PathExists judge whether file or directory exists.
return:
	true, nil : exist.
	false, nil: not exist.
	false, err: I don't know.
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// TryMkDir try to mkdir strDirectory if strDirectory is not exists
func TryMkDir(strDirectory string) {
	if exists, _ := PathExists(strDirectory); !exists {
		if err := os.MkdirAll(strDirectory, 0777); err != nil {
			log.Error("mkdir directory = %v failed, err = %v", strDirectory, err)
		}
	}
}

// HTTPRsp HTTP response
func HTTPRsp(w http.ResponseWriter, iHTTPCode int, strMessage string) {
	w.WriteHeader(iHTTPCode)
	if len(strMessage) > 0 {
		w.Write([]byte(strMessage))
	}
}
