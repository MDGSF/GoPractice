package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/MDGSF/utils/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gDB *gorm.DB

var deviceID string = "2bd40a1228c57f7b"

func main() {

	config, err := ioutil.ReadFile("config")
	if err != nil {
		log.Error("failed to read config, err = %v", err)
		return
	}
	configstr := strings.TrimSpace(string(config))

	gDB, err = gorm.Open("mysql", configstr)
	if err != nil {
		log.Error("failed to open mysql, err = %v", err)
		return
	}
	defer gDB.Close()
	log.Info("open mysql success")

	queryAdas()
	queryDms()
	queryStatus()
}

func queryAdas() {
	for i := 1; i <= 18; i++ {
		queryOneDayAdas(i)
	}
}

func queryOneDayAdas(day int) {
	tableName := fmt.Sprintf("adas_events_202011%+02v", day)

	curDB := gDB.Table(tableName).Select("*")
	curDB = curDB.Where("device = ?", deviceID)

	var count int
	curDB.Count(&count)

	daystr := fmt.Sprintf("2020_11_%+02v", day)
	fmt.Printf("adas event, %v, count = %v\n", daystr, count)
}

func queryDms() {
	for i := 1; i <= 18; i++ {
		queryOneDayDms(i)
	}
}

func queryOneDayDms(day int) {
	tableName := fmt.Sprintf("dms_events_202011%+02v", day)

	curDB := gDB.Table(tableName).Select("*")
	curDB = curDB.Where("device = ?", deviceID)

	var count int
	curDB.Count(&count)

	daystr := fmt.Sprintf("2020_11_%+02v", day)
	fmt.Printf("dms event, %v, count = %v\n", daystr, count)
}

func queryStatus() {
	for i := 1; i <= 18; i++ {
		queryOneDayStatus(i)
	}
}

func queryOneDayStatus(day int) {
	tableName := fmt.Sprintf("adas_status_202011%+02v", day)

	curDB := gDB.Table(tableName).Select("*")
	curDB = curDB.Where("device = ?", deviceID)

	var count int
	curDB.Count(&count)

	daystr := fmt.Sprintf("2020_11_%+02v", day)
	fmt.Printf("status, %v, count = %v\n", daystr, count)
}
