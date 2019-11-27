package main

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/MDGSF/utils/log"
)

var execFileName = "/tmp/test"

func main() {
	binaryBytes, err := hex.DecodeString(BinaryHello)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = ioutil.WriteFile(execFileName, binaryBytes, 0777)
	if err != nil {
		log.Fatalf("%v", err)
	}

	cmd := exec.Command(execFileName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
