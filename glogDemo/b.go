package main

import "github.com/golang/glog"

func fb() {
	glog.V(8).Info("b.go level 8 info")
	glog.V(9).Info("b.go level 9 info")
}
