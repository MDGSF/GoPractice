package main

/*
#cgo CFLAGS : -I.
#cgo LDFLAGS: -L. -lcamera
#cgo LDFLAGS: -L/root/git/thirdparty_repo/libusb-1.0/android_armv8 -lusb1.0
#cgo LDFLAGS: -L/usr/local/ndk-toolchain-arm64/sysroot/usr/lib -lm -lc -lz -llog -lstdc++

#include "camera.h"
*/
import "C"

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	"unsafe"

	"github.com/MDGSF/utils/log"
)

func main() {
	fmt.Println("camera demo starting...")

	C.startCamera(1, 1000, 0, 1000, 704, 576, 25, 10)

	fmt.Println("camera demo started.")

	for {
		time.Sleep(time.Second * 10)

		//getAdasImage() //ok
		getAdasVideo()

		//getDsmImage()

		fmt.Println("----------------------------------------------------")
	}
}

func getAdasVideo() {
	var iVideoSize C.int
	var pcVideoData *C.char
	var cret C.int

	curTime := time.Now().Unix()
	var startTime C.long = C.long(curTime - 20)
	var endTime C.long = C.long(curTime)
	cret = C.getAdasVideo(startTime, endTime, &iVideoSize, &pcVideoData)
	defer C.free(unsafe.Pointer(pcVideoData))

	ret := int(cret)
	videoSize := int(iVideoSize)
	videoData := C.GoBytes(unsafe.Pointer(pcVideoData), iVideoSize)

	if ret != 0 {
		log.Error("getAdasVideo failed, ret = %v", ret)
		return
	}

	log.Info("getAdasVideo success, VideoSize = %v\n", videoSize)

	err := ioutil.WriteFile(strconv.Itoa(int(time.Now().Unix()))+".mp4", videoData, 0666)
	if err != nil {
		log.Fatal("write file failed, err = %v", err)
	}
}

func getAdasImage() {
	var iImageSize C.int
	var pcImageData *C.char
	var cret C.int
	cret = C.getLatestAdasImage(&iImageSize, &pcImageData)
	defer C.free(unsafe.Pointer(pcImageData))

	ret := int(cret)
	imageSize := int(iImageSize)
	imageData := C.GoBytes(unsafe.Pointer(pcImageData), iImageSize)

	if ret != 0 {
		log.Error("getLatestAdasImage failed, ret = %v", ret)
		return
	}

	log.Info("getLatestAdasImage success, imageSize = %v\n", imageSize)

	err := ioutil.WriteFile(strconv.Itoa(int(time.Now().Unix()))+".jpg", imageData, 0666)
	if err != nil {
		log.Fatal("write file failed, err = %v", err)
	}
}

func getDsmImage() {
	var iImageSize C.int
	var pcImageData *C.char
	var cret C.int
	cret = C.getLatestDsmImage(&iImageSize, &pcImageData)
	defer C.free(unsafe.Pointer(pcImageData))

	ret := int(cret)
	imageSize := int(iImageSize)
	imageData := C.GoBytes(unsafe.Pointer(pcImageData), iImageSize)

	if ret != 0 {
		log.Error("getLatestDsmImage failed, ret = %v", ret)
		return
	}

	log.Info("getLatestDsmImage success, imageSize = %v\n", imageSize)

	err := ioutil.WriteFile(strconv.Itoa(int(time.Now().Unix()))+".jpg", imageData, 0666)
	if err != nil {
		log.Fatal("write file failed, err = %v", err)
	}
}
