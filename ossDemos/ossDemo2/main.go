package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/MDGSF/utils/log/mwriter"
)

var gStart time.Time

func processImage(startTime time.Time, days int) {
	start := time.Now()
	log.Info("[%v] processImage start = %v", startTime, start)
	defer func() {
		elapsed := time.Since(start)
		log.Info("[%v] processImage end, time elapsed = %v", startTime, elapsed)
	}()

	tryTimes := 0

	for i := 0; i < days; i++ {

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
			log.Error("curTimeStr = %v, err = %v, output = %v",
				curTimeStr, err, string(output))

			if tryTimes > 3 {
				tryTimes = 0
			} else {
				tryTimes += 1
				i -= 1
			}

			time.Sleep(time.Second)

			continue
		}
		log.Info("%v", string(output))
		log.Info("for end, forElapsed = %v, gElapsed = %v",
			time.Since(forStart), time.Since(gStart))
	}
}

func processVideo(startTime time.Time, days int) {
	start := time.Now()
	log.Info("[%v] processVideo start = %v", startTime, start)
	defer func() {
		elapsed := time.Since(start)
		log.Info("[%v] processVideo end, time elapsed = %v", startTime, elapsed)
	}()

	tryTimes := 0

	for i := 0; i < days; i++ {

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

			if tryTimes > 3 {
				tryTimes = 0
			} else {
				tryTimes += 1
				i -= 1
			}

			time.Sleep(time.Second)

			continue
		}
		log.Info("%v", string(output))
		log.Info("for end, forElapsed = %v, gElapsed = %v",
			time.Since(forStart), time.Since(gStart))
	}
}

func process(name string, startTime time.Time, days int) {
	start := time.Now()
	log.Info("%v start = %v", name, start)
	defer func() {
		elapsed := time.Since(start)
		log.Info("%v end, time elapsed = %v", name, elapsed)
	}()

	processImage(startTime, days)
	processVideo(startTime, days)
}

func process2018() {
	startTime := time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	process("process2018", startTime, 365)
}

func process2019() {
	startTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
	process("process2019", startTime, 365)
}

func process2020() {
	// startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	// process("process2020", startTime, 304) // left 2 months

	startTime := time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local)
	process("process2020", startTime, 61)
}

func process_start_to_end(name string, startYear, startMonth, startDay,
	endYear, endMonth, endDay int) {
	startTime := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.Local)
	endTime := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.Local)
	elapsed := endTime.Sub(startTime)
	days := int(elapsed.Hours())/24 + 1
	process(name, startTime, days)
}

func process2021() {
	process_start_to_end("2021", 2021, 1, 1, 2021, 3, 31)
}

var gStartYear int
var gStartMonth int
var gStartDay int
var gEndYear int
var gEndMonth int
var gEndDay int

func main() {

	flag.IntVar(&gStartYear, "StartYear", 2021, "StartYear")
	flag.IntVar(&gStartMonth, "StartMonth", 1, "StartMonth")
	flag.IntVar(&gStartDay, "StartDay", 1, "StartDay")

	flag.IntVar(&gEndYear, "EndYear", 2021, "EndYear")
	flag.IntVar(&gEndMonth, "EndMonth", 1, "EndMonth")
	flag.IntVar(&gEndDay, "EndDay", 1, "EndDay")

	flag.Parse()

	w := mwriter.New("/home/huangjian/oss_process.log",
		10*1024*1024,
		time.Duration(4320*int64(time.Minute)))
	log.SetOutput(w)
	log.SetLevel(log.NameToLevel("info"))
	log.SetIsTerminal(log.NotTerminal)

	gStart = time.Now()
	log.Info("main gStart = %v", gStart)
	defer func() {
		elapsed := time.Since(gStart)
		log.Info("main end, time elapsed = %v", elapsed)
	}()

	//process2018()
	//process2019()
	//process2020()
	//process2021()

	process_start_to_end("test", gStartYear, gStartMonth, gStartDay,
		gEndYear, gEndMonth, gEndDay)
}
