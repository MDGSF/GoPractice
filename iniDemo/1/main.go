package main

import (
	"github.com/astaxie/beego/config"
)

func main() {
	cnf, err := config.NewConfig("ini", "conf.ini")
	if err != nil {
		return
	}

	if err := cnf.SaveConfigFile("conf2.ini"); err != nil {
		return
	}
}
