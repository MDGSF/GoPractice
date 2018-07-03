
all:
	CC=/usr/local/ndk-toolchain-arm64/bin/aarch64-linux-android-gcc \
	GOOS=android \
	GOARCH=arm64 \
	CGO_ENABLED=1 \
	GODEBUG=cgocheck=2 \
	go build -o camera-arm64 main.go

copy:
	cp /root/git/lms_m4/camera.h .
	cp /root/git/lms_m4/arm/lib/libcamera.so .

push:
	adb push camera-arm64 /sdcard/run
	adb push libcamera.so /sdcard/run

clean:
	-rm -rf camera-arm64

