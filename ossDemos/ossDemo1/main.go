package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/MDGSF/utils/log/mwriter"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const Version = "1.0.0"

const Usage = `Usage: oss_process [Options]

Options:
  --help                  print this help message and exit
  --version               print the version and exit
  --config FILE           set the configuration file
`

var ConfigFile string
var BuildTime string

var MainConfig TConfig

var ossClient *oss.Client

var gAllFileCounter *big.Int
var gAllfileSizeInBytes *big.Int

var gAllAdasFileCounter *big.Int
var gAllAdasFileSizeInBytes *big.Int

var gAllDmsFileCounter *big.Int
var gAllDmsFileSizeInBytes *big.Int

var mutex sync.Mutex

func init() {
	gAllFileCounter = big.NewInt(0)
	gAllfileSizeInBytes = big.NewInt(0)
	gAllAdasFileCounter = big.NewInt(0)
	gAllAdasFileSizeInBytes = big.NewInt(0)
	gAllDmsFileCounter = big.NewInt(0)
	gAllDmsFileSizeInBytes = big.NewInt(0)
}

// TConfig 配置文件
type TConfig struct {
	// TerminalLog 终端日志配置
	TerminalLog TLogConfig `json:"terminallog"`

	// FileLog 文件日志配置，如果 TerminalLog 和 FileLog 同时打开，那么只有 TerminalLog 生效。
	FileLog TLogConfig `json:"filelog"`

	OssConfig TOssConfig `json:"oss"`
}

// TLogConfig 日志相关配置
type TLogConfig struct {

	// Enable 是否开启日志
	Enable bool `json:"enable"`

	// Level 日志等级： fatal error warn info debug verbose
	Level string `json:"level"`

	// FileName 文件日志的文件名
	FileName string `json:"filename"`

	FileMaxSize int `json:"FileMaxSize"`

	FileTTLMinute int64 `json:"FileTTLMinute"`
}

// TOssConfig OSS 配置
type TOssConfig struct {
	EndPoint           string `json:"endpoint"`
	AccessKeyID        string `json:"AccessKeyID"`
	KeySECRET          string `json:"KeySECRET"`
	BucketName         string `json:"BucketName"`
	PageNumber         int    `json:"PageNumber"`
	PrintElapsedNumber int64  `json:"PrintElapsedNumber"`
}

func loadConfigFile(strConfigFileName string) error {
	if len(strConfigFileName) <= 0 {
		return fmt.Errorf("empty config file")
	}

	aucFileContent, err := ioutil.ReadFile(strConfigFileName)
	if err != nil {
		log.Warn("failed to read config file [%v]: %s\n", strConfigFileName, err)
		return err
	}

	if err := json.Unmarshal(aucFileContent, &MainConfig); err != nil {
		log.Warn("invalid json file [%v]: %s", strConfigFileName, err)
		return err
	}

	return nil
}

func initLog() {
	if MainConfig.TerminalLog.Enable {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.NameToLevel(MainConfig.TerminalLog.Level))
		log.SetIsTerminal(log.IsTerminal)
	} else if MainConfig.FileLog.Enable {
		w := mwriter.New(MainConfig.FileLog.FileName,
			MainConfig.FileLog.FileMaxSize,
			time.Duration(MainConfig.FileLog.FileTTLMinute*int64(time.Minute)))
		log.SetOutput(w)
		log.SetLevel(log.NameToLevel(MainConfig.FileLog.Level))
		log.SetIsTerminal(log.NotTerminal)
		log.Info("MainConfig.FileLog.Enable = %v", MainConfig.FileLog.Enable)
		log.Info("MainConfig.FileLog.FileName = %v", MainConfig.FileLog.FileName)
		log.Info("MainConfig.FileLog.FileMaxSize = %v bytes", MainConfig.FileLog.FileMaxSize)
		log.Info("MainConfig.FileLog.FileTTLMinute = %v minutes",
			MainConfig.FileLog.FileTTLMinute)
	}
}

func process() {
	// 统计一个半月之内的数据

	wg := &sync.WaitGroup{}

	currentTime := time.Now()
	for i := -45; i <= 0; i++ {
		oldTime := currentTime.AddDate(0, 0, i)
		dateStr := fmt.Sprintf("%04v-%02v-%02v", oldTime.Year(), int(oldTime.Month()),
			oldTime.Day())

		wg.Add(1)
		go list_file(wg, dateStr)
	}

	wg.Wait()
}

func list_file(wg *sync.WaitGroup, dateStr string) {
	defer func() {
		wg.Done()
	}()

	fmt.Println()
	log.Info("list_file start, dateStr = %v", dateStr)

	var fileCounter int64 = 0
	var fileSizeInBytes int64 = 0

	var adasFileCounter int64 = 0
	var adasFileSizeInBytes int64 = 0

	var dmsFileCounter int64 = 0
	var dmsFileSizeInBytes int64 = 0

	// 创建OSSClient实例。
	ossClient, err := oss.New(MainConfig.OssConfig.EndPoint,
		MainConfig.OssConfig.AccessKeyID, MainConfig.OssConfig.KeySECRET)
	if err != nil {
		log.Error("Error: %v", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := ossClient.Bucket(MainConfig.OssConfig.BucketName)
	if err != nil {
		log.Error("Error: %v", err)
		os.Exit(-1)
	}

	// 分页列举包含指定前缀的文件。每页列举80个。
	prefix := oss.Prefix(dateStr)
	marker := oss.Marker("")
	for {
		lsRes, err := bucket.ListObjects(oss.MaxKeys(MainConfig.OssConfig.PageNumber), marker, prefix)
		if err != nil {
			log.Error("Error: %v", err)
			time.Sleep(time.Second)
			continue
		}

		prefix = oss.Prefix(lsRes.Prefix)
		marker = oss.Marker(lsRes.NextMarker)

		for _, object := range lsRes.Objects {
			log.Verbose("object.XMLName: %#v", object.XMLName)
			log.Verbose("object.Key: %#v", object.Key)
			log.Verbose("object.Type: %#v", object.Type)
			log.Verbose("object.Size: %#v", object.Size)
			log.Verbose("object.ETag: %#v", object.ETag)
			log.Verbose("object.Owner: %#v", object.Owner)
			log.Verbose("object.LastModified: %#v", object.LastModified)
			log.Verbose("object.StorageClass: %#v", object.StorageClass)

			if object.Size == 0 {
				continue
			}
			fileCounter += 1
			fileSizeInBytes += object.Size

			if strings.Contains(object.Key, "_adas.mp4") {
				adasFileCounter += 1
				adasFileSizeInBytes += object.Size
			}

			if strings.Contains(object.Key, "_dms.mp4") {
				dmsFileCounter += 1
				dmsFileSizeInBytes += object.Size
			}

			if fileCounter%MainConfig.OssConfig.PrintElapsedNumber == 0 {
				mutex.Lock()
				log.Info("[%v] fileCounter = %v", dateStr, fileCounter)
				log.Info("[%v] fileSizeInBytes = %v", dateStr, fileSizeInBytes)
				log.Info("[%v] adasFileCounter = %v", dateStr, adasFileCounter)
				log.Info("[%v] adasFileSizeInBytes = %v", dateStr, adasFileSizeInBytes)
				log.Info("[%v] dmsFileCounter = %v", dateStr, dmsFileCounter)
				log.Info("[%v] dmsFileSizeInBytes = %v\n", dateStr, dmsFileSizeInBytes)
				mutex.Unlock()
			}
		}
		log.Verbose("len(Objects): %v", len(lsRes.Objects))

		if !lsRes.IsTruncated {
			break
		}
	}

	mutex.Lock()
	log.Info("list_file end, dateStr = %v", dateStr)
	log.Info("fileCounter = %v", fileCounter)
	log.Info("fileSizeInBytes = %v", fileSizeInBytes)
	log.Info("adasFileCounter = %v", adasFileCounter)
	log.Info("adasFileSizeInBytes = %v", adasFileSizeInBytes)
	log.Info("dmsFileCounter = %v", dmsFileCounter)
	log.Info("dmsFileSizeInBytes = %v\n", dmsFileSizeInBytes)

	gAllFileCounter.Add(gAllFileCounter, big.NewInt(fileCounter))
	gAllfileSizeInBytes.Add(gAllfileSizeInBytes, big.NewInt(fileSizeInBytes))

	gAllAdasFileCounter.Add(gAllAdasFileCounter, big.NewInt(adasFileCounter))
	gAllAdasFileSizeInBytes.Add(gAllAdasFileSizeInBytes, big.NewInt(adasFileSizeInBytes))

	gAllDmsFileCounter.Add(gAllDmsFileCounter, big.NewInt(dmsFileCounter))
	gAllDmsFileSizeInBytes.Add(gAllDmsFileSizeInBytes, big.NewInt(dmsFileSizeInBytes))

	log.Info("gAllFileCounter = %v", gAllFileCounter)
	log.Info("gAllfileSizeInBytes = %v", gAllfileSizeInBytes)
	log.Info("gAllAdasFileCounter = %v", gAllAdasFileCounter)
	log.Info("gAllAdasFileSizeInBytes = %v", gAllAdasFileSizeInBytes)
	log.Info("gAllDmsFileCounter = %v", gAllDmsFileCounter)
	log.Info("gAllDmsFileSizeInBytes = %v\n", gAllDmsFileSizeInBytes)

	mutex.Unlock()
}

func main() {
	help := flag.Bool("help", false, "print this help message and exit")
	version := flag.Bool("version", false, "print the version and exit")
	flag.StringVar(&ConfigFile, "config", "config.json", "set the configuration file")

	flag.Parse()

	if *help {
		fmt.Println(Usage)
		return
	}

	if *version {
		fmt.Printf("ossp %s (%s) [%s-%s] (%s)\n", Version, BuildTime,
			runtime.GOOS, runtime.GOARCH, runtime.Version())
		return
	}

	if err := loadConfigFile(ConfigFile); err != nil {
		log.Error("err = %v", err)
		return
	}

	initLog()

	process()

	fmt.Println()
	log.Info("gAllFileCounter = %v", gAllFileCounter)
	log.Info("gAllfileSizeInBytes = %v", gAllfileSizeInBytes)
	log.Info("gAllAdasFileCounter = %v", gAllAdasFileCounter)
	log.Info("gAllAdasFileSizeInBytes = %v", gAllAdasFileSizeInBytes)
	log.Info("gAllDmsFileCounter = %v", gAllDmsFileCounter)
	log.Info("gAllDmsFileSizeInBytes = %v", gAllDmsFileSizeInBytes)
}
