ANDROIDGCC=/usr/local/ndk-toolchain-arm64/bin/aarch64-linux-android-gcc

all: targetAndroid

targetAndroid:
	CC=/usr/local/ndk-toolchain-arm64/bin/aarch64-linux-android-gcc \
	GOOS=android \
	GOARCH=arm64 \
	CGO_ENABLED=1 \
	go build -o main-arm64 main.go
	mkdir ../bin/android
	mv main-arm64 ../bin/android
	cp ../cppLib/android/dynamicLib/libadd.so ../bin/android/

