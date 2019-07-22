/*
flowout -sub "output.car_info.speed" -save "output.car_info.speed" -save-dir "/data/rem_data"

flowout -sub "output.car_info.speed" -save "output.car_info.speed" -save-dir "/data/rem_data" -limit 1

flowout -sub "debug.hub.lane,debug.hub.tsr,output.car_info.gps,output.car_info.speed" -save "debug.hub.lane,debug.hub.tsr,output.car_info.gps,output.car_info.speed" -save-dir "/data/rem_data"
*/

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/MDGSF/utils/log"
	"github.com/vmihailenco/msgpack"
)

// StringToInt64 convert string to int64
func StringToInt64(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}

// StringToFloat32 convert string to float32
func StringToFloat32(str string) float32 {
	val, _ := strconv.ParseFloat(str, 32)
	return float32(val)
}

// StringToFloat64 convert string to float64
func StringToFloat64(str string) float64 {
	val, _ := strconv.ParseFloat(str, 64)
	return float64(val)
}

func strToBytes(str string) []byte {
	i, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		log.Error("err = %v", err)
		return nil
	}
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, i)
	bytes := bytesBuffer.Bytes()
	return bytes
}

func getMicroTimeFromFileName(filename string) (int64, error) {
	baseFileName := filepath.Base(filename)
	arr1 := strings.Split(baseFileName, "-")
	if len(arr1) != 3 {
		log.Error("arr1 = %v", arr1)
		return 0, fmt.Errorf("invalid filename = %v", baseFileName)
	}

	arr2 := strings.Split(arr1[0], ".")
	if len(arr2) != 2 {
		log.Error("arr2 = %v", arr2)
		return 0, fmt.Errorf("invalid filename = %v", baseFileName)
	}

	return strconv.ParseInt(arr2[1], 10, 64)
}

func processLane(srcFile, dstFile string) {
	log.Info("process lane file = %v", srcFile)

	result := make(map[string]interface{})
	result["type"] = "lane"

	data, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}

	m := make(map[string]interface{})
	err = msgpack.Unmarshal(data, &m)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}

	srcLaneLinesI, ok := m["lanelines"]
	if !ok {
		log.Error("no lanelines exists")
		return
	}

	srcLaneLines, ok := srcLaneLinesI.([]interface{})
	if !ok {
		log.Error("invalid lanelines type")
		return
	}

	alllanelines := make([]interface{}, 0)
	for k, oneLaneLineI := range srcLaneLines {

		oneLaneLine, ok := oneLaneLineI.(map[string]interface{})
		if !ok {
			log.Error("invalid oneLaneLineI type")
			return
		}

		dstOneLaneLine := make(map[string]interface{})
		dstOneLaneLine["id"] = k
		dstOneLaneLine["pts"] = oneLaneLine["bird_view_pts"]
		dstOneLaneLine["lane_type"] = oneLaneLine["type"]
		dstOneLaneLine["color"] = oneLaneLine["color"]
		dstOneLaneLine["width"] = oneLaneLine["width"]

		alllanelines = append(alllanelines, dstOneLaneLine)
	}
	result["lanelines"] = alllanelines

	t, err := getMicroTimeFromFileName(srcFile)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}
	result["ts"] = t

	resultdata, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}
	ioutil.WriteFile(dstFile, resultdata, 0644)
}

func processTsr(srcFile, dstFile string) {
	log.Info("process tsr file = %v", srcFile)
}

func processGPS(srcFile, dstFile string) {
	log.Info("process gps file = %v", srcFile)

	data, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}

	srcGPS := &TMsgCarGPS{}
	srcGPS.ParseMsgCarGPS(data)

	result := make(map[string]interface{})
	result["type"] = "gps"

	t, err := getMicroTimeFromFileName(srcFile)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}
	result["ts"] = t

	pos := make([]float64, 2)
	pos[0] = srcGPS.Latitude
	pos[1] = srcGPS.Longitude
	result["pos"] = pos

	result["alt"] = srcGPS.Altitude
	result["speed"] = srcGPS.Speed
	result["bearing"] = srcGPS.Bearing
	result["accuracy"] = srcGPS.Accuracy

	resultdata, err := json.Marshal(result)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}
	ioutil.WriteFile(dstFile, resultdata, 0644)
}

type TMsgCarGPS struct {
	Index     int64   `msgpack:"index"`
	SysTime   int64   `msgpack:"systime"`
	Flags     uint16  `msgpack:"flags"`
	Latitude  float64 `msgpack:"latitude"`
	Longitude float64 `msgpack:"longitude"`
	Altitude  float64 `msgpack:"altitude"`
	Speed     float32 `msgpack:"speed"`
	Bearing   float32 `msgpack:"bearing"`
	Accuracy  float32 `msgpack:"accuracy"`
	Time      int64   `msgpack:"timestamp"`
}

func (t *TMsgCarGPS) ParseMsgCarGPS(buf []byte) {
	if err := msgpack.Unmarshal(buf, t); err == nil {
		return
	}
	str := string(buf)
	arr := strings.Split(str, " ")
	if len(arr) == 8 {
		valB := strToBytes(arr[0])
		var tmp int32
		bytesBuffer := bytes.NewBuffer(valB)
		binary.Read(bytesBuffer, binary.LittleEndian, &tmp)

		t.Flags = uint16(tmp)
		t.Latitude = StringToFloat64(arr[1])
		t.Longitude = StringToFloat64(arr[2])
		t.Altitude = StringToFloat64(arr[3])
		t.Speed = StringToFloat32(arr[4])
		t.Bearing = StringToFloat32(arr[5])
		t.Accuracy = StringToFloat32(arr[6])
		t.Time = StringToInt64(arr[7])
		return
	}

	if len(arr) == 10 {
		valB := strToBytes(arr[2])
		var tmp int32
		bytesBuffer := bytes.NewBuffer(valB)
		binary.Read(bytesBuffer, binary.LittleEndian, &tmp)

		t.Index = StringToInt64(arr[1])
		t.Flags = uint16(tmp)
		t.Latitude = StringToFloat64(arr[3])
		t.Longitude = StringToFloat64(arr[4])
		t.Altitude = StringToFloat64(arr[5])
		t.Speed = StringToFloat32(arr[6])
		t.Bearing = StringToFloat32(arr[7])
		t.Accuracy = StringToFloat32(arr[8])
		t.Time = StringToInt64(arr[9])
		return
	}
	t = nil
}

func processSpeed(srcFile, dstFile string) {
	log.Info("process speed file = %v", srcFile)

	data, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}

	speed := StringToFloat32(string(data))

	result := make(map[string]interface{})
	result["type"] = "car_speed"

	t, err := getMicroTimeFromFileName(srcFile)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}
	result["ts"] = t

	result["speed"] = speed

	resultdata, err := json.Marshal(result)
	if err != nil {
		log.Error("err = %v, srcFile = %v", err, srcFile)
		return
	}
	ioutil.WriteFile(dstFile, resultdata, 0644)
}

func processOneFile(srcFile, dstFile string) {
	if strings.HasSuffix(srcFile, "debug.hub.lane.dat") {
		processLane(srcFile, dstFile)
	} else if strings.HasSuffix(srcFile, "debug.hub.tsr.dat") {
		processTsr(srcFile, dstFile)
	} else if strings.HasSuffix(srcFile, "output.car_info.gps.dat") {
		processGPS(srcFile, dstFile)
	} else if strings.HasSuffix(srcFile, "output.car_info.speed.dat") {
		processSpeed(srcFile, dstFile)
	} else {
		log.Info("Unknown file = %v", srcFile)
	}
}

func processRawData(srcDir, dstDir string) {
	os.MkdirAll(dstDir, 0755)

	entries, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatal("err = ", err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		srcFile := filepath.Join(srcDir, entry.Name())
		dstFile := filepath.Join(dstDir, entry.Name())
		processOneFile(srcFile, dstFile)
	}
}

func main() {
	srcdir := "/home/huangjian/a/local/gopath/src/github.com/MDGSF/GoPractice/remdata/rem_data"
	dstdir := "/home/huangjian/a/local/gopath/src/github.com/MDGSF/GoPractice/remdata/rem_output"
	processRawData(srcdir, dstdir)
}
