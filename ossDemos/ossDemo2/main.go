package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/MDGSF/utils/log"
)

var gStart time.Time

func processImage(startTime time.Time) {
	for i := 0; i < 365; i++ {

		forStart := time.Now()

		curTime := startTime.AddDate(0, 0, i)
		curTimeStr := fmt.Sprintf("%+02v-%+02v-%+02v", curTime.Year(),
			int(curTime.Month()), curTime.Day())

		//ossutil rm oss://minieyeimage/2018-xx-xx/ -r -f --include "*.jpg" --exclude "*_dms.jpg"
		name := "/home/huangjian/ossutil64"
		args := make([]string, 0)
		args = append(args, "rm")
		args = append(args, fmt.Sprintf("oss://minieyeimage/%v/", curTimeStr))
		args = append(args, "-r")
		args = append(args, "-f")
		args = append(args, "--include")
		args = append(args, "*.jpg")
		args = append(args, "--exclude")
		args = append(args, "*_dms.jpg")

		log.Info("args = %v", args)

		cmd := exec.Command(name, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Error("curTimeStr = %v, err = %v", curTimeStr, err)
			break
		}
		log.Info("%v", string(output))
		log.Info("for end, forElapsed = %v, gElapsed = %v",
			time.Since(forStart), time.Since(gStart))
	}
}

func processVideo(startTime time.Time) {
	for i := 0; i < 365; i++ {

		forStart := time.Now()

		curTime := startTime.AddDate(0, 0, i)
		curTimeStr := fmt.Sprintf("%+02v-%+02v-%+02v", curTime.Year(),
			int(curTime.Month()), curTime.Day())

		//ossutil rm oss://minieyeimage/2018-xx-xx/ -r -f --include "*.jpg" --exclude "*_dms.jpg"
		name := "/home/huangjian/ossutil64"
		args := make([]string, 0)
		args = append(args, "rm")
		args = append(args, fmt.Sprintf("oss://minieyevideo/%v/", curTimeStr))
		args = append(args, "-r")
		args = append(args, "-f")
		args = append(args, "--include")
		args = append(args, "*.mp4")
		args = append(args, "--exclude")
		args = append(args, "*_dms.mp4")

		log.Info("args = %v", args)

		cmd := exec.Command(name, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Error("curTimeStr = %v, err = %v", curTimeStr, err)
			break
		}
		log.Info("%v", string(output))
		log.Info("for end, forElapsed = %v, gElapsed = %v",
			time.Since(forStart), time.Since(gStart))
	}
}

func process2018() {
	startTime := time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	//processImage(startTime)
	processVideo(startTime)
}

func process2019() {
	startTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
	processImage(startTime)
}

func main() {
	gStart = time.Now()
	defer func() {
		elapsed := time.Since(gStart)
		log.Info("time used: %v", elapsed)
	}()

	process2018()
}
