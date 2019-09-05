/*
https://cloud.minieye.cc/download/devices/00270d94822e70cb/dates/2019-09-05/events_v2.csv
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MDGSF/utils/log"
	"github.com/astaxie/bat/httplib"
)

// mediaType minieyevideo, minieyeimage
var mediaType = "minieyevideo"

var outputDir = "event_output"

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

func getCookies() []*http.Cookie {
	data, err := ioutil.ReadFile("cookies.json")
	if err != nil {
		log.Fatalf("err = %v", err)
	}

	type TCookies struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Path   string `json:"path"`
		Domain string `json:"domain"`
	}

	type TCookiesJSON struct {
		Cookies []TCookies
	}

	cookies := &TCookiesJSON{}
	err = json.Unmarshal(data, cookies)
	if err != nil {
		log.Fatalf("err = %v", err)
	}

	result := make([]*http.Cookie, 0)
	for k := range cookies.Cookies {
		oneCookie := &http.Cookie{}

		oneCookie.Name = cookies.Cookies[k].Name
		oneCookie.Value = cookies.Cookies[k].Value
		oneCookie.Path = cookies.Cookies[k].Path
		oneCookie.Domain = cookies.Cookies[k].Domain

		result = append(result, oneCookie)
	}
	return result
}

func main() {

	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	log.Info("curDir = %v", curDir)

	eventOutputDir := filepath.Join(curDir, outputDir)
	os.MkdirAll(eventOutputDir, 0755)

	cookies := getCookies()

	divicesList := getList("devices.md")
	dateList := getList("datetime.md")

	log.Info("dateList = %v", dateList)
	log.Info("divicesList = %v", divicesList)

	for i := range dateList {
		for j := range divicesList {
			rawURL := fmt.Sprintf("https://cloud.minieye.cc/download/devices/%v/dates/%v/events_v2.csv",
				divicesList[j], dateList[i],
			)

			filename := fmt.Sprintf("%v-events-%v_v2.csv", divicesList[j], dateList[i])
			filePathName := filepath.Join(curDir, outputDir, filename)
			log.Info("curDir = %v", curDir)
			log.Info("outputDir = %v", outputDir)
			log.Info("filename = %v", filename)
			log.Info("filePathName = %v", filePathName)

			req := httplib.Get(rawURL)
			req.SetEnableCookie(true)

			for k := range cookies {
				req.SetCookie(cookies[k])
			}

			err = req.ToFile(filePathName)
			if err != nil {
				log.Fatalf("err = %v", err)
			}
		}
	}

	//err = ucmd.PackTarGz(curDir, outputDir)
	//if err != nil {
	//	log.Fatalf("err = %v", err)
	//}

	//os.RemoveAll(outputDir)
}
