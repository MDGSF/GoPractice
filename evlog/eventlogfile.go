package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/MDGSF/utils/log"
	"go.uber.org/zap"
)

// TEventLogProcessor 保存上传的事件到文件中，用来上传到 oss
type TEventLogProcessor struct {
	EventLogFileChan chan []byte
	maxsize          int64 //byte
	preYear          int
	preMonth         time.Month
	preDay           int
	preIndex         int
	logDir           string
	preFileName      string
	f                *os.File
}

// CreateEventLogProcessor 创建一个事件日志处理器
func CreateEventLogProcessor(logDir string,
	EventLogFileChan chan []byte,
	maxsize int64) *TEventLogProcessor {
	p := &TEventLogProcessor{}
	p.EventLogFileChan = EventLogFileChan
	p.maxsize = maxsize
	p.preIndex = 1
	p.logDir = logDir
	p.updateDate()
	p.createLogFile()
	return p
}

// Start 启动事件日志处理器
func (p *TEventLogProcessor) Start() {
	go p.run()
}

func (p *TEventLogProcessor) run() {
	for {
		select {
		case msg, ok := <-p.EventLogFileChan:
			if !ok {
				log.Error("EventLogFileChan not ok")
				return
			}

			log.Info("msg start", time.Now().UnixNano())

			if p.needDailyRotate() {
				p.rotateDaily()
			} else if p.needRotateSize() {
				p.rotateSize()
			}

			p.writeMsg(msg)

			log.Info("msg end", time.Now().UnixNano())
		}
	}
}

func (p *TEventLogProcessor) needDailyRotate() bool {
	year, month, day := time.Now().Date()
	if year != p.preYear || month != p.preMonth || day != p.preDay {
		return true
	}
	return false
}

func (p *TEventLogProcessor) updateDate() {
	p.preYear, p.preMonth, p.preDay = time.Now().Date()
}

func (p *TEventLogProcessor) rotateDaily() {
	var err error

	if p.f != nil {
		err = p.f.Close()
		p.f = nil
		if err != nil {
			return
		}
	}

	p.updateDate()
	p.createLogFile()
}

func (p *TEventLogProcessor) needRotateSize() bool {
	if p.curFileSize() > p.maxsize {
		return true
	}
	return false
}

func (p *TEventLogProcessor) rotateSize() {
	var err error

	if p.f != nil {
		err = p.f.Close()
		p.f = nil
		if err != nil {
			return
		}
	}

	p.preIndex++
	p.createLogFile()
}

func (p *TEventLogProcessor) curFileSize() int64 {
	fileInfo, _ := p.f.Stat()
	return fileInfo.Size()
}

func (p *TEventLogProcessor) createLogFile() {
	var err error
	p.preFileName = fmt.Sprintf("%04d-%02d-%02d_%04d.json",
		p.preYear, p.preMonth, p.preDay, p.preIndex)
	fileName := filepath.Join(p.logDir, p.preFileName)
	p.f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Error("TEventLogProcessor create log file failed",
			zap.Any("err", err))
		return
	}
}

func (p *TEventLogProcessor) writeMsg(msg []byte) {
	if _, err := p.f.Write(msg); err != nil {
		log.Error("TEventLogProcessor write msg failed",
			zap.Any("err", err))
	}
}
