package main

import "github.com/golang/glog"

func fa() {
	glog.V(5).Info("a.go level 5 info")
	glog.V(6).Info("a.go level 6 info")
}
