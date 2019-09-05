/*
https://cloud.minieye.cc/download/devices/00270d94822e70cb/dates/2019-09-05/events_v2.csv
*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MDGSF/utils"
	"github.com/MDGSF/utils/log"
	"github.com/MDGSF/utils/ucmd"
)

// mediaType minieyevideo, minieyeimage
var mediaType = "minieyevideo"

var outputDir = "event_output"
var eventOutputDir string

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

	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	log.Info("curDir = %v", curDir)

	eventOutputDir = filepath.Join(curDir, outputDir)
	os.MkdirAll(eventOutputDir, 0755)

	runtimeDir := filepath.Join(curDir, "runtime")
	os.MkdirAll(runtimeDir, 0755)

	scriptTplBytes, err := ioutil.ReadFile("curltemplate.sh")
	if err != nil {
		log.Fatalf("err = %v", err)
	}

	re, _ := regexp.Compile("https://cloud.minieye.cc/download/devices/(.*)/dates/(.*)/events_v2.csv")
	subs := re.FindSubmatch(scriptTplBytes)
	for k := range subs {
		log.Info("subs[%v] = %v", k, string(subs[k]))
	}
	if len(subs) != 3 {
		log.Fatalf("len(subs) = %v", len(subs))
	}

	divicesList := getList("devices.md")
	dateList := getList("datetime.md")

	log.Info("dateList = %v", dateList)
	log.Info("divicesList = %v", divicesList)

	for i := range dateList {
		for j := range divicesList {

			curScript := bytes.Replace(scriptTplBytes, subs[1], []byte(divicesList[j]), 1)
			curScript = bytes.Replace(curScript, subs[2], []byte(dateList[i]), 1)
			curScript = bytes.Replace(curScript, []byte("curl "), []byte("curl -s "), 1)

			scriptFileName := fmt.Sprintf("%v_%v.sh", dateList[i], divicesList[j])
			scriptPathName := filepath.Join(runtimeDir, scriptFileName)
			ioutil.WriteFile(scriptPathName, curScript, 0777)
		}
	}

	RunScripts(runtimeDir)

	err = ucmd.PackTarGz(curDir, outputDir)
	if err != nil {
		log.Fatalf("err = %v", err)
	}

	os.RemoveAll(outputDir)
	os.RemoveAll(runtimeDir)
}

func RunScripts(scriptDir string) error {
	if exist, _ := utils.PathExists(scriptDir); exist {

		entries, err := ioutil.ReadDir(scriptDir)
		if err != nil {
			log.Error("err = %v", err)
			return err
		}

		log.Info("len(entries) = %v", len(entries))

		for k := range entries {
			entry := entries[k]
			if entry.IsDir() {
				continue
			}
			filePathName := filepath.Join(scriptDir, entry.Name())
			cmd := exec.Command("bash", "-c", filePathName)
			if output, err := cmd.CombinedOutput(); err != nil {
				log.Error(fmt.Sprintf("err = %v, filePathName = %v", err, filePathName))
			} else {
				log.Info("exec script %v success", entry.Name())
				fileName := entry.Name()
				parts := strings.Split(fileName, ".")
				newFileName := fmt.Sprintf("%v.csv", parts[0])
				newFilePathName := filepath.Join(eventOutputDir, newFileName)
				ioutil.WriteFile(newFilePathName, output, 0644)
			}
		}
	}

	return nil
}
