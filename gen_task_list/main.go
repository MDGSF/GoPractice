package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

var VideoDir string
var TaskFileName string

type Task struct {
	Repeat     int    `json:"repeat"`
	DetectFlag string `json:"detect_flag"`
	Name       string `json:"name"`
	Source     string `json:"source"`
	Path       string `json:"path"`
}

type TaskList struct {
	Tasks []Task `json:"data"`
}

var gTaskCount int = 0
var gTaskList TaskList

func main() {
	flag.StringVar(&VideoDir, "video", "", "set the video dir")
	flag.StringVar(&TaskFileName, "task", "task.json", "set the task file name")

	flag.Parse()

	gTaskList.Tasks = make([]Task, 0)

	walkDir(VideoDir)

	data, _ := json.MarshalIndent(&gTaskList, "", "  ")
	ioutil.WriteFile(TaskFileName, data, 0644)
}

func walkDir(directory string) {
	entries, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			entryFullPath := filepath.Join(directory, entry.Name())
			if entry.Name() == "video" {
				gTaskCount++
				task := Task{
					Repeat:     1,
					DetectFlag: filepath.Join(entryFullPath, "base.flag"),
					Name:       "task" + strconv.Itoa(gTaskCount),
					Source:     entryFullPath,
					Path:       filepath.Join(entryFullPath, "log.txt"),
				}
				gTaskList.Tasks = append(gTaskList.Tasks, task)
			} else {
				walkDir(entryFullPath)
			}
		}
	}
}
