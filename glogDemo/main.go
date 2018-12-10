package main

import (
	"flag"

	"github.com/golang/glog"
)

/*
go run main.go a.go b.go -log_dir="./" -alsologtostderr -v=3 -vmodule=a*=5,b*=7

go run main.go a.go b.go -log_dir="./" -alsologtostderr -v=3 -vmodule=a*=5,b*=8

go run main.go a.go b.go -log_dir="./" -alsologtostderr -v=3 -vmodule=a*=5,b*=9

go run main.go a.go b.go -log_dir="./" -alsologtostderr -vmodule=a*=5,b*=9

-v=3
是设置所有 glog.V(level) 中 level <=3 的都执行

-vmodule=a*=5
是设置所有以 a 开头的 go 文件 level <=5 的都执行
*/

func main() {
	flag.Parse()

	glog.Info("info log")
	glog.Warning("warning log")
	glog.Error("error log")

	glog.V(1).Info("level 1 info")
	glog.V(2).Info("level 2 info")

	fa()
	fb()

	glog.Flush()
}
