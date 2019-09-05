/*
ossutil cp oss://minieyevideo/2019-09-01/09b64617414c4837 oss_output -r -f -u --include "*_adas.mp4"
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/MDGSF/utils/log"
	"github.com/MDGSF/utils/ucmd"
)

// mediaType minieyevideo, minieyeimage
var mediaType = "minieyevideo"

var outputDir = "oss_output"

var fileNameFilter = "*_adas.mp4"

func getList(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("err = %v", err)
	}

	rawlist := strings.Split(string(data), "\n")
	result := make([]string, 0)
	for k := range rawlist {
		cur := rawlist[k]
		cur = strings.TrimSpace(cur)
		if len(cur) == 0 {
			continue
		}

		result = append(result, cur)
	}
	return result
}

func main() {
	divicesList := getList("devices.md")
	dateList := getList("datetime.md")

	log.Info("dateList = %v", dateList)
	log.Info("divicesList = %v", divicesList)

	for i := range dateList {
		for j := range divicesList {
			sourceURL := fmt.Sprintf("oss://%v/%v/%v",
				mediaType, dateList[i], divicesList[j])

			argsList := []string{
				"cp", sourceURL, outputDir, "-r", "-f", "-u",
			}

			if len(fileNameFilter) > 0 {
				argsList = append(argsList, "--include", fileNameFilter)
			}

			log.Info("ossutil %v", argsList)
			cmd := exec.Command("/root/ossutil64", argsList...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("err = %v", err)
			}
			log.Info("out = %v", string(out))
		}
	}

	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	log.Info("curDir = %v", curDir)

	err = ucmd.PackTarGz(curDir, outputDir)
	if err != nil {
		log.Fatalf("err = %v", err)
	}

	os.RemoveAll(outputDir)
}
