package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	test2()
}

func test3() {
	err := DecompressFile("test.zip", ".")
	if err != nil {
		fmt.Println(err)
	}
}

func test2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		z := NewZipWithWriter(w)
		defer z.Close()

		err := z.CompressOneFile("a.txt", "image/b.txt")
		if err != nil {
			log.Fatal(err)
		}

		err = z.CompressData([]byte("Happy Day"), "h.txt")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Disposition", "attachment; filename="+"test.zip")
		w.Header().Add("Content-Description", "File Transfer")
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Content-Transfer-Encoding", "binary")
		w.Header().Add("Expires", "0")
		w.Header().Add("Cache-Control", "must-revalidate")
		w.Header().Add("Pragma", "public")

	})

	http.ListenAndServe("127.0.0.1:8000", nil)
}

func test1() {
	z := NewZipWithFile("test.zip")
	defer z.Close()

	err := z.CompressOneFile("a.txt", "image/b.txt")
	if err != nil {
		log.Fatal(err)
	}
}

type Zip struct {
	writer *zip.Writer
}

func NewZipWithFile(fileName string) *Zip {
	z := &Zip{}
	d, err := os.Create(fileName)
	if err != nil {
		return nil
	}

	z.writer = zip.NewWriter(d)
	return z
}

func NewZipWithWriter(writer io.Writer) *Zip {
	z := &Zip{}
	z.writer = zip.NewWriter(writer)
	return z
}

func (z *Zip) Close() {
	z.writer.Close()
}

func (z *Zip) CompressOneFile(srcFileName, dstFileName string) error {
	f, err := os.OpenFile(srcFileName, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = dstFileName

	writer, err := z.writer.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, f)
	if err != nil {
		return err
	}

	return nil
}

func (z *Zip) CompressData(data []byte, dstFileName string) error {
	writer, err := z.writer.Create(dstFileName)
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func DecompressFile(zipFileName, dst string) error {
	reader, err := zip.OpenReader(zipFileName)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		log.Printf("fileName = %v", file.Name)
		path := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}

		targetFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(targetFile, rc)
		if err != nil {
			return err
		}

		rc.Close()
		targetFile.Close()
	}

	return nil
}
