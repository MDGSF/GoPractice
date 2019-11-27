package main

import (
	"encoding/hex"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/MDGSF/utils/log"
)

var binaryFileName = "hello"
var outputFileName = "binaryHello.go"

func main() {
	content, err := ioutil.ReadFile(binaryFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}

	dst := make([]byte, hex.EncodedLen(len(content)))
	hex.Encode(dst, content)
	contentStr := string(dst)

	t := template.New("generate binary demo")
	t, _ = t.Parse(`
package main

var BinaryHello = "{{.Content}}"
	`)

	type TGenerateBinary struct {
		Content string
	}
	p := TGenerateBinary{Content: contentStr}

	outputFile, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer outputFile.Close()

	err = t.Execute(outputFile, p)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Info("generate %v success", outputFileName)
}
