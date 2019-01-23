package main

import (
	"archive/zip"
	"compress/flate"
	"fmt"

	"github.com/mholt/archiver"
)

func main() {
	z := archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		OverwriteExisting:      true,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ImplicitTopLevelFolder: false,
		ContinueOnError:        false,
	}

	err := z.Archive([]string{"test/testdata", "test/file.txt"}, "./1/test.zip")
	if err != nil {
		fmt.Println(err)
	}

	err = z.Walk("./1/test.zip", func(f archiver.File) error {
		zfh, ok := f.Header.(zip.FileHeader)
		if ok {
			fmt.Println("Filename:", zfh.Name)
		}
		return nil
	})

}
