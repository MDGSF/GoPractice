package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type TLocation struct {
	Offset uint32
	Size   uint32
}

var ChunkFile *os.File

var m map[uint64]TLocation

var RangKeyGenerator uint64

func main() {
	f, err := os.OpenFile("chunk", os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("open chunk failed, err = ", err)
	}
	ChunkFile = f

	fileInfo, err := ChunkFile.Stat()
	if err != nil {
		log.Fatal("stat failed, err = ", err)
	}
	log.Println(fileInfo.Name(), fileInfo.Size())

	m = make(map[uint64]TLocation)

	r := mux.NewRouter()
	r.HandleFunc("/add", Add)
	r.HandleFunc("/get", Get)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	log.Println("add")

	r.ParseMultipartForm(32 << 20)
	files := r.MultipartForm.File["uploadfile"]
	for _, fileHeader := range files {

		randkey := uint64(RangKeyGenerator)
		RangKeyGenerator++

		fileInfo, _ := ChunkFile.Stat()

		offset := fileInfo.Size()

		file, err := fileHeader.Open()
		if err != nil {
			log.Println(err)
			continue
		}

		written, err := io.Copy(ChunkFile, file)
		if err != nil {
			log.Println(err)
			continue
		}

		if written != fileHeader.Size {
			log.Println(written, fileHeader.Size)
			continue
		}

		m[randkey] = TLocation{
			Offset: uint32(offset),
			Size:   uint32(written),
		}

		log.Println(randkey, offset, written)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	log.Println("get")

	r.ParseForm()

	log.Println(r.Form, r.PostForm)

	randkeystr := r.Form.Get("key")
	randkey, _ := strconv.Atoi(randkeystr)

	location, ok := m[uint64(randkey)]
	if !ok {
		log.Println("invalid randkey = ", randkey)
		return
	}

	b := make([]byte, location.Size)
	n, err := ChunkFile.ReadAt(b, int64(location.Offset))
	if err != nil {
		log.Println(err)
		return
	}

	if uint32(n) != location.Size {
		log.Println(n, location.Size)
		return
	}

	buffer := bytes.NewBuffer(b)

	io.Copy(w, buffer)
}
